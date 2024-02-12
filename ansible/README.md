# Documentation: Installation Processes

## Introduction
This documentation provides a detailed guide to resolving issues encountered during the installation processes, specifically focusing on conflicts and configuration adjustments related to starting the containerd process and disabling swap in the system.

## Issues
During the installation processes, several issues were identified:

### 1. Conflict Preventing Containerd Process Start
Observation: The containerd process encountered conflicts and failed to start properly.

Resolution:
To resolve this issue, it was necessary to configure containerd to enable System Cgroup. This configuration was added to the /etc/containerd/config.toml file.

```toml
version = 2
[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
          [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
            SystemdCgroup = true
```

### 2. Swap Re-enablement After System Restart
Observation: Despite disabling swap from the operating system, it was noted that after restarting the machine, the swap was enabled again.

Resolution:
After investigating the issue, it was discovered that the ZRam service needed to be disabled to prevent swap re-enablement. The following command was used to check for ZRam service:

```shell
sudo systemctl list-units --type=service | grep -i zram
```

## Conclusion

By following the provided resolutions, the conflicts preventing the proper startup of the containerd process were resolved, and the issue of swap re-enablement after system restart was mitigated. These adjustments ensure smoother installation processes and maintain system stability.