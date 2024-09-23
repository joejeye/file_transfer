# File Transfer Between Devices

This is a simple file transfer application that allows you to transfer files between devices on the same network. 

## Features

- Transfer files between devices on the same network
- Easy to use
- Lightweight

## Installation
Clone the repository

## Configuration
Create a `config.toml` file in the `server` directory with the following:
```
download_directory = "<directory_to_download_files_to>"
listen_port = "<port_to_listen_on>"
```
where `<directory_to_download_files_to>` is the directory you want to download the files to, and `<port_to_listen_on>` is the port you want the server to listen on.

## Running the server
```
go run server/main.go
```

## Sending a file using the client
```
go run client/main.go <server_ip>:<server_port> <path_to_file>
```

where `<server_ip>` is the IP address of the server, `<server_port>` is the port the server is listening on, and `<path_to_file>` is the file you want to send.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.
