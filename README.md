# File Transfer Between Devices

This is a simple file transfer application that allows you to transfer files between devices on LAN. 

## Features

- Transfer files between devices within the same LAN
- Automatically detects available receivers
- Easy to use
- Lightweight

## Prerequisite
Golang

## Installation
Clone the repository

## Configuration
Create a `config.toml` file in the `server` directory with the following:
```
download_directory = "<directory_to_download_files_to>"
listen_port = "<port_to_listen_on>"
```
where `<directory_to_download_files_to>` is the directory you want to 
download the files to, and `<port_to_listen_on>` is the port you want the
server to listen on. If the file path contains backslashes, make sure to 
escape them or use raw string literals.

## Running the server
Run the following command at the server-side:
```
go run file_transfer_naive/server
```

## Sending a file using the client
Run the following command at the client-side:
```
go run file_transfer_naive/client
```
and follow the instructions. The client will first search the LAN for available receivers, 
then prompt you to select a server to send the file to. Lastly, it will ask you to enter 
the path to the file you want to send.

## License
This project is licensed under the MIT License. See the `License.txt` file for more details.
