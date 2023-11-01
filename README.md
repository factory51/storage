# storagesolutions.io

## Intro
This is a small example application to show Mysql Master <-> Master replication used in a realistic scenario. The application works as a semi-secured distributed key/value store, with the intention of allowing users to store data securely and remotely.


## Servers
 The application is running on 3 stock Debian 12 VPS instances.

### storage-01: 167.172.41.43
This instance is a combined application and database server.

**Configuration**
- go1.21.3
- 10.11.4-MariaDB-1~deb12u1-log [Replication source storage-02]

### storage-02: 188.166.10.136
This instance is also a combined application and database server.

**Configuration**
- go1.21.3
- 10.11.4-MariaDB-1~deb12u1-log [Replication source storage-01]



### storage-balancer: 178.128.250.86

This instance is purely setup as a load balancer to distribute traffic between the two application / database instances.

**Configuration**
- nginx/1.22.1 only

## Usage

The example application has been configured to show which application server and therefore copy of the replicated database has been used in the request, eg: - 

```
% curl -H "authorization:13a90fab880a7e82209fb5c8b86ac125" -i "http://storage.storagesolutions.io/get/test"     
HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Wed, 01 Nov 2023 11:08:48 GMT
Content-Type: application/json
Content-Length: 100
Connection: keep-alive
X-Clacks-Overhead: GNU Terry Pratchett
X-Powered-By: Factory51
X-Server-Ident: storage-01
```

```
% curl -H "authorization:13a90fab880a7e82209fb5c8b86ac125" -i "http://storage.storagesolutions.io/get/test"     
HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Wed, 01 Nov 2023 11:08:48 GMT
Content-Type: application/json
Content-Length: 100
Connection: keep-alive
X-Clacks-Overhead: GNU Terry Pratchett
X-Powered-By: Factory51
X-Server-Ident: storage-02
```

As you can see from the two identical requests the `X-Server-Ident` header differs. The load balancer is splitting which application instance is handling the request.

## Usage Examples

**note**: All requests require the authorization header set to `13a90fab880a7e82209fb5c8b86ac125` as this is functioning as basic access control to the application.


### Get

The application handles 2 workflows for Get. One where the `key` is supplied, and the full corresponding entry is returned from the database, and one where a `key` isn't supplied and the application returns all avaliable `key` values that will return a meaningful response.

```
% curl -H "authorization:13a90fab880a7e82209fb5c8b86ac125" -i "http://storage.storagesolutions.io/get/test"
HTTP/1.1 200 OK
Server: nginx/1.22.1
Date: Wed, 01 Nov 2023 11:20:29 GMT
Content-Type: application/json
Content-Length: 100
Connection: keep-alive
X-Clacks-Overhead: GNU Terry Pratchett
X-Powered-By: Factory51
X-Server-Ident: storage-01

{"created_at":"2023-11-01 08:10:33","client_ident":"storage-02","key":"test","value":"Hello World!"}
```