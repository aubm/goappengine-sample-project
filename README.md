This is a sample project built to illustrate the deployment of Go app on AppEngine standard.

As I write, a working solution is to use the `goapp` command, as follow:

```
goapp deploy -application <my-project> -version beta-001 app
```

One solution that does not work for now, is to use the `gcloud` command, as follow:

```
gcloud --project=<my-project> app deploy --version=beta-001 app/app.yaml
```

The command fails and produces the following output:

```
You are about to deploy the following services:
 - <my-project>/default/beta-001 (from [/home/<user>/gopath/src/github.com/aubm/goappengine-sample-project/app/app.yaml])
     Deploying to URL: [https://<my-project>.appspot.com]
Do you want to continue (Y/n)?  
Beginning deployment of service [default]...
Copying files to Google Cloud Storage...
Synchronizing files to [staging.<my-project>.appspot.com].
File upload done.
Updating service [default]...failed.                                                                                                                                                              
ERROR: (gcloud.app.deploy) Error Response: [9] Deployment contains files that cannot be compiled: Compile failed:
2016/12/02 03:16:51 go-app-builder: build timing: 1×compile (106ms total), 0×link (0s total)
2016/12/02 03:16:51 go-app-builder: failed running compile: exit status 1
app.go:6: can't find import: "github.com/aubm/goappengine-sample-project/api"
```
