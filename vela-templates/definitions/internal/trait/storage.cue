storage: {
	type: "trait"
	annotations: {}
	labels: {}
	description: "Add storages on K8s pod for your workload which follows the pod spec in path 'spec.template'."
	attributes: {
		appliesToWorkloads: ["deployments.apps"]
		podDisruptive: true
	}
}
template: {
	pvcVolumesList: *[
			for v in parameter.pvc {
			{
				name: "pvc-" + v.name
				persistentVolumeClaim: claimName: v.name
			}
		},
	] | []

	configMapVolumesList: *[
				for v in parameter.configMap {
			{
				name: "configmap-" + v.name
				configMap: {
					defaultMode: v.defaultMode
					name:        v.name
					if v.items != _|_ {
						items: v.items
					}
				}
			}
		},
	] | []

	secretVolumesList: *[
				for v in parameter.secret {
			{
				name: "secret-" + v.name
				secret: {
					defaultMode: v.defaultMode
					secretName:  v.name
					if v.items != _|_ {
						items: v.items
					}
				}
			}
		},
	] | []

	emptyDirVolumesList: *[
				for v in parameter.emptyDir {
			{
				name: "emptydir-" + v.name
				emptyDir: {
					medium: v.medium
				}
			}
		},
	] | []

	pvcVolumeMountsList: *[
				for v in parameter.pvc {
			if v.volumeMode == "Filesystem" {
				{
					name:      "pvc-" + v.name
					mountPath: v.mountPath
				}
			}
		},
	] | []

	configMapVolumeMountsList: *[
					for v in parameter.configMap {
			{
				name:      "configmap-" + v.name
				mountPath: v.mountPath
			}
		},
	] | []

	configMapEnvMountsList: *[
				for v in parameter.configMap if v.mountToEnv != _|_ {
			{
				name: v.mountToEnv.envName
				valueFrom: configMapKeyRef: {
					name: v.name
					key:  v.mountToEnv.configMapKey
				}
			}
		},
	] | []

	secretVolumeMountsList: *[
				for v in parameter.secret {
			{
				name:      "secret-" + v.name
				mountPath: v.mountPath
			}
		},
	] | []

	secretEnvMountsList: *[
				for v in parameter.secret if v.mountToEnv != _|_ {
			{
				name: v.mountToEnv.envName
				valueFrom: secretKeyRef: {
					name: v.name
					key:  v.mountToEnv.secretKey
				}
			}
		},
	] | []

	emptyDirVolumeMountsList: *[
					for v in parameter.emptyDir {
			{
				name:      "emptydir-" + v.name
				mountPath: v.mountPath
			}
		},
	] | []

	volumeDevicesList: *[
				for v in parameter.pvc if v.volumeMode == "Block" {
			{
				name:       "pvc-" + v.name
				devicePath: v.mountPath
			}
		},
	] | []

	patch: spec: template: spec: {
		// +patchKey=name
		volumes: pvcVolumesList + configMapVolumesList + secretVolumesList + emptyDirVolumesList

		containers: [...{
			// +patchKey=name
			env: configMapEnvMountsList + secretEnvMountsList
			// +patchKey=name
			volumeDevices: volumeDevicesList
			// +patchKey=name
			volumeMounts: pvcVolumeMountsList + configMapVolumeMountsList + secretVolumeMountsList + emptyDirVolumeMountsList
		}]

	}

	outputs: {
		for v in parameter.pvc {
			if v.mountOnly == false {
				"pvc-\(v.name)": {
					apiVersion: "v1"
					kind:       "PersistentVolumeClaim"
					metadata: {
						name: v.name
					}
					spec: {
						accessModes: v.accessModes
						volumeMode:  v.volumeMode
						if v.volumeName != _|_ {
							volumeName: v.volumeName
						}
						if v.storageClassName != _|_ {
							storageClassName: v.storageClassName
						}

						if v.resources.requests.storage == _|_ {
							resources: requests: storage: "1Gi"
						}
						if v.resources.requests.storage != _|_ {
							resources: requests: storage: v.resources.requests.storage
						}
						if v.resources.limits.storage != _|_ {
							resources: limits: storage: v.resources.limits.storage
						}
						if v.dataSourceRef != _|_ {
							dataSourceRef: v.dataSourceRef
						}
						if v.dataSource != _|_ {
							dataSource: v.dataSource
						}
						if v.selector != _|_ {
							dataSource: v.selector
						}
					}
				}
			}
		}

		for v in parameter.configMap {
			if v.mountOnly == false {
				"configmap-\(v.name)": {
					apiVersion: "v1"
					kind:       "ConfigMap"
					metadata: name: v.name
					if v.data != _|_ {
						data: v.data
					}
				}
			}
		}

		for v in parameter.secret {
			if v.mountOnly == false {
				"secret-\(v.name)": {
					apiVersion: "v1"
					kind:       "Secret"
					metadata: name: v.name
					if v.data != _|_ {
						data: v.data
					}
					if v.stringData != _|_ {
						stringData: v.stringData
					}
				}
			}
		}

	}

	parameter: {
		// +usage=Declare pvc type storage
		pvc?: [...{
			name:              string
			mountOnly:         *false | bool
			mountPath:         string
			volumeMode:        *"Filesystem" | string
			volumeName?:       string
			accessModes:       *["ReadWriteOnce"] | [...string]
			storageClassName?: string
			resources?: {
				requests: storage: =~"^([1-9][0-9]{0,63})(E|P|T|G|M|K|Ei|Pi|Ti|Gi|Mi|Ki)$"
				limits?: storage:  =~"^([1-9][0-9]{0,63})(E|P|T|G|M|K|Ei|Pi|Ti|Gi|Mi|Ki)$"
			}
			dataSourceRef?: {
				name:     string
				kind:     string
				apiGroup: string
			}
			dataSource?: {
				name:     string
				kind:     string
				apiGroup: string
			}
			selector?: {
				matchLabels?: [string]: string
				matchExpressions?: {
					key: string
					values: [...string]
					operator: string
				}
			}
		}]

		// +usage=Declare config map type storage
		configMap?: [...{
			name:      string
			mountOnly: *false | bool
			mountToEnv?: {
				envName:      string
				configMapKey: string
			}
			mountPath:   string
			defaultMode: *420 | int
			readOnly:    *false | bool
			data?: {...}
			items?: [...{
				key:  string
				path: string
				mode: *511 | int
			}]
		}]

		// +usage=Declare secret type storage
		secret?: [...{
			name:      string
			mountOnly: *false | bool
			mountToEnv?: {
				envName:   string
				secretKey: string
			}
			mountPath:   string
			defaultMode: *420 | int
			readOnly:    *false | bool
			stringData?: {...}
			data?: {...}
			items?: [...{
				key:  string
				path: string
				mode: *511 | int
			}]
		}]

		// +usage=Declare empty dir type storage
		emptyDir?: [...{
			name:      string
			mountPath: string
			medium:    *"" | "Memory"
		}]
	}

}
