<h1>PowerTop Monitoring</h1>
<h3>Using Prometheus and Graphana</h3>

<p>PowerTOP is a terminal-based diagnosis tool that helps you to monitor power usage by programs running on a Linux system when it is not plugged on to a power source , which makes it suitable for unreliable power sources
For PowerTop to work in Edge Devices easily , this a image was to be required , which is build with help of Docker image and is available at <a href="https://hub.docker.com/">DockerHub registry</a>. </p>
<p>Furthermore the stats can be acquired with the help of Prometheus metrics , and can be stored in internal Prometheus TSDB . These data can be really helpful for alert management or even a visual representation of the stats using tools like grafana etc</p>

<p>While running thousands of application in edge devices the monitoring and optimisation of power consumption is crucial </p>

<h3>Local SetUp</h3>
<h4>Pre-requisite</h4>
<ol>
   <li>Linux Environment<ul>
  </ul></li>
  <li><a href="https://docs.docker.com/compose/install/">Docker Compose </a><ul>
  </ul></li>
  <li><a href="https://podman.io/getting-started/installation.html#installing-on-linux">Podman</a><ul>
  </ul></li>
</ol>


<h3>Dev SetUp</h3>

open up a terminal

Restart podman  <code>sudo systemctl restart podman</code>   
Enter the folder  <code>cd powertop_monitoring</code>  
Run using Go , requires super user priviledge  <code>sudo go run ./main.go</code>  
Check the metrics using curl   <code>curl http://localhost:8886/metrics</code>  
