> **Note**
> This project has been merged into the instance manager which can be found [here](https://github.com/dhis2-sre/im-manager)

# Instance Database Manager

## Quick start
```sh
export ACCESS_TOKEN && eval (./login.sh) && echo $ACCESS_TOKEN
wget https://databases.dhis2.org/sierra-leone/2.36.0/dhis2-db-sierra-leone.sql.gz -P ~/Downloads
./upload.sh whoami ~/Downloads/dhis2-db-sierra-leone.sql.gz
```


# Token valid by key found in jwks.json

```sh
export ACCESS_TOKEN=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3OTYzODQwNzcsImlhdCI6MTY0MDY5MTQ3NywidXNlciI6eyJJRCI6NCwiQ3JlYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiVXBkYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiRGVsZXRlZEF0IjpudWxsLCJFbWFpbCI6InNvbWVvbmVAc29tZXRoaW5nLm9yZyIsIkdyb3VwcyI6bnVsbCwiQWRtaW5Hcm91cHMiOm51bGx9fQ.FnPIu36kV1T-Jix5Wy-HsZeqxQI6Q_7HQ14C1DWKHETIBSk-vLQ_sCMHVPKA42utEDFI3Xpmf6Gyzv9aPU_Cvg-JDazRprfrZBqn4LzSmT6K3HGoKoQ0b5G8exxz0Ote8NQDB1NBZmYvD1gpVVisCvzaewJTRAvRA3DS0n_O4kU5QENdLNPfWFo0rXOC83sLBsEIe2Ce4TiRrepOCSQKE-_rwQQSA3w30MhFmhAY7Ozcd9i69mtfcvqjORdNJ-zREgiw8B2g9oh7byE1h2oxjvoKC3WRfPeSYoRY6GuMHSSWJdzFKIswlZHdWU1GicPJASBbkKGbP5n5O6FXyeo0bw
```


# TODO

*         - name: data
          emptyDir:
            sizeLimit: {{ .Values.dataSizeLimit }}
dataSizeLimit seems to have no effect and the mounted disk doesn't seem to be a volume... But it is writable

* Timeout... Eventually we'll need to run the export etc. in a background thread. Should we just return 202 or an id or something else?
* Don't implement the endpoint on the database manager. Rather implement /instances/:id/save-as on the manager and let that produce an event which will be consumed by the database manager
* Introduce some sort of status property on the database model... dumped, zipped... ready?
* Store the dump log on the database model as well. The output from `dump.EnableVerbose()` and... ? gzip output?
* 


# TODO

*         - name: data
          emptyDir:
            sizeLimit: {{ .Values.dataSizeLimit }}
dataSizeLimit seems to have no effect and the mounted disk doesn't seem to be a volume... But it is writable

* Timeout... Eventually we'll need to run the export etc. in a background thread. Should we just return 202 or an id or something else?
* Don't implement the endpoint on the database manager. Rather implement /instances/:id/save-as on the manager and let that produce an event which will be consumed by the database manager
* Introduce some sort of status property on the database model... dumped, zipped... ready?
* Store the dump log on the database model as well. The output from `dump.EnableVerbose()` and... ? gzip output?
