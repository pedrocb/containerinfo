docker_build('containerinfo', '.',
    dockerfile='Dockerfile')
k8s_yaml(listdir('k8s'))
k8s_resource('containerinfo', port_forwards=8000)
