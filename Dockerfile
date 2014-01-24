from ubuntu:12.04
maintainer evan hazlett "<ehazlett@arcus.io>"
run apt-get update
run apt-get install -y ca-certificates
add slacker-wiki /usr/local/bin/slacker-wiki
add run.sh /usr/local/bin/run
expose 8080
cmd ["/usr/local/bin/run"]
