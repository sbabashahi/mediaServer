**AIR**   ->        https://github.com/cosmtrek/air
    after installation of it in project add alias air='./bin/air' to .bashrc and add .air.toml to project.


docker
    docker build -t mediamanager .

    sudo docker run --rm -it -p 8080:8080 -v /home/mastisa/go/src/mediamanager/user/:/go/src/mediamanager/user/ --name mediaserver my-golang-app:v3