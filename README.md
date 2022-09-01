<h1>PowerTop Monitoring</h1>
<h3>Using Prometheus and Graphana</h3>

<p>PowerTOP is a terminal-based diagnosis tool that helps you to monitor power usage by programs running on a Linux system when it is not plugged on to a power source , which makes it suitable for unreliable power sources
For PowerTop to work in Edge Devices easily , this a image was to be required , which is build with help of Docker image and is available at <a href="https://hub.docker.com/">DockerHub registry</a>. </p>
<p>Furthermore the stats can be acquired with the help of Prometheus metrics , and can be stored in internal Prometheus TSDB . These data can be really helpful for alert management or even a visual representation of the stats using tools like grafana etc</p>

<p>While running thousands of application in edge devices the monitoring and optimisation of power consumption is crucial </p>

<h3>Local SetUp</h3>
<h4>Pre-requisite</h4>
<ol>
   <li>Linux Environment for running without container<ul>
</ol>


<h3>Dev SetUp</h3>

for this powertop is needed to be pre installed

open up a terminal

1. clone the repo 

2. go in the folder <code>cd powertop_monitoring</code>  

3. run using go compiler <code>sudo go run cmd/main.go</code>  
   powertop requires sudo permission to access the system stats

4.bare prometheus metrics can be seen using <code>curl 0.0.0.0:8887/metrics</code>

<h3>Running Using Docker</h3>

1.for this you used need --priviledge flag , which would give it access to host energy stats
    <code> docker run -p 8887:8887--privileged powertopcsv:v2</code>  
2.bare prometheus metrics can be seen using <code> curl 0.0.0.0:8887/metrics</code>  

These can be run with graphana and prometheus easily with the docker compose file


