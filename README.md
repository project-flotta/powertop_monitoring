<h1>PowerTop Monitoring</h1>


<p>PowerTOP is a terminal-based diagnosis tool that helps you to monitor power usage by programs running on a Linux system when it is not plugged on to a power source , which makes it suitable for unreliable power sources
For PowerTop to work in Edge Devices easily , this a image was to be required , which is build with help of Docker image and is available at <a href="https://hub.docker.com/">DockerHub registry</a>. </p>
<p>Furthermore the stats can be acquired with the help of Prometheus metrics , and can be stored in internal Prometheus TSDB . These data can be really helpful for alert management or even a visual representation of the stats using tools like grafana etc</p>

<p>While running thousands of application in edge devices the monitoring and optimisation of power consumption is crucial </p>

For installation of flotta-operator and flotta-edge device  
follow the flotta guide

[kind installation](https://project-flotta.io/documentation/v0_2_0/gsg/kind.html)

[flotta-dev-cli](https://project-flotta.io/flotta/2022/07/20/developer-cli.html)

<h3>Deploying PowerTop Workload</h3>

The powertop monitoring application would be deployed as workloads.
Details on how to deploying workloads are in 
[flotta workloads deployment](https://project-flotta.io/documentation/v0_2_0/gsg/running_workloads.html)

The yaml for the workload  :-

```yaml
apiVersion: management.project-flotta.io/v1alpha1
kind: EdgeWorkload
metadata:
  name: powertop
spec:
  metrics:
    interval: 5
    path: "/metrics"
    port: 8887
  deviceSelector:
    matchLabels:
      app: foo
  type: pod
  pod:
    spec:
      containers:
        - name: powertop
          image: docker.io/sibseh/powertopcsv:v2
          securityContext:
            privileged: true
          volumeMounts:
            - mountPath: /lib/modules
              name: lib-modules
            - mountPath: /sys/kernel/
              name: tracing
            - mountPath: /usr/src/
              name: usr
          ports :
            - containerPort: 8887
              hostPort: 8887
      volumes:
        - hostPath:
          path: /lib/modules
          type: Directory
          name: lib-modules
        - hostPath:
          path: /sys/kernel
          type: Directory
          name: tracing
        - hostPath:
          path: /usr/src/
          type: Directory
          name: usr
```

<h3>Monitoring Using Thanos</h3>

A thanos and graphana set up can be used for monitoring visually 

More details can be found in :-

[flotta observability](https://project-flotta.io/documentation/latest/operations/observability.html)

[writing-metrics-to-control-plane](https://project-flotta.io/flotta/2022/04/11/writing-metrics-to-control-plane.html
)


For Thanos receiver 

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: thanos-receiver
  labels:
    app: thanos-receiver
spec:
  replicas: 1
  selector:
    matchLabels:
      app: thanos-receiver
  template:
    metadata:
      labels:
        app: thanos-receiver
    spec:
      containers:
        - name: receive
          image: quay.io/thanos/thanos:v0.24.0
          command:
            - /bin/thanos
            - receive
            - --log.level
            - debug
            - --label
            - "receiver=\"0\""
            - --remote-write.address
            - 0.0.0.0:10908
        - name: query
          image: quay.io/thanos/thanos:v0.24.0
          command:
            - /bin/thanos
            - query
            - --log.level
            - debug
            - --http-address
            - 0.0.0.0:9090
            - --grpc-address
            - 0.0.0.0:11901
            - --endpoint
            - 127.0.0.1:10901
---
apiVersion: v1
kind: Service
metadata:
  name: thanos-receiver
spec:
  type: NodePort
  selector:
    app: thanos-receiver
  ports:
    - port: 80
      targetPort: 10908
      nodePort: 30030
      name: endpoint
    - port: 9090
      targetPort: 9090
      name: admin
      
---
apiVersion: v1
kind: Service
metadata:
  name: thanos-receiver
spec:
  type: NodePort
  selector:
    app: thanos-receiver
  ports:
    - port: 80
      targetPort: 10908
      nodePort: 30030
      name: endpoint
    - port: 9090
      targetPort: 9090
      name: admin

```


For Graphana
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grafana
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      name: grafana
      labels:
        app: grafana
    spec:
      containers:
        - name: grafana
          image: grafana/grafana:latest
          ports:
            - name: grafana
              containerPort: 3000
          resources:
            limits:
              memory: "1Gi"
              cpu: "1000m"
            requests:
              memory: 500M
              cpu: "500m"
          volumeMounts:
            - mountPath: /var/lib/grafana
              name: grafana-storage
      volumes:
        - name: grafana-storage

---

View the complete the documentation [here](https://docs.google.com/document/d/1COQ66hhWg9gm_kQUjO1IGN5ERfKr-3eXLq5dcOhR1yY/edit?usp=sharing)


