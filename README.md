# HubbleWebAssignment

Prerequisites

1.Install Docker 
    https://docs.docker.com/engine/install/

2.clone the project
    https://github.com/BalamuruganP25/HubbleWeb.git

3.Run the project 
  3.1 Go inside the project 
        cd HubbleWeb
  3.2 Build project 
        docker build -t hubbleweb .
  3.3 Run the project
      docker container run -p 8080:8080 -it hubbleweb:latest

    Note - Client & server in same project