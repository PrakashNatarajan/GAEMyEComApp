# GAEMyEComApp
Run The Main(Default) Service.
`/home/prakash/GoogleCloudSdk/bin/dev_appserver.py default/app.yaml`

Run The Order Service.
`/home/prakash/GoogleCloudSdk/bin/dev_appserver.py orders/orders.yaml`

Important:
Data will be created in the local system when running the app in local.

Deploy The App in Google App Engine.
`gcloud app deploy`

`gcloud app deploy service1/service1.yaml --verbosity=info`
`gcloud app deploy service2/service2.yaml --verbosity=info`
`gcloud app deploy default/app.yaml --verbosity=info`
`gcloud app deploy dispatch.yaml --verbosity=info`

List the App Versions.
`gcloud app versions list`

Delete the Particular Version from na App.
`gcloud app versions delete <VersionNo>`
`gcloud app versions delete 20170402t173941`

Open the App in the Browser after Deployment.
`gcloud app browse`

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)
 