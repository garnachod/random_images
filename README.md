## random_images
[![<ORG_NAME>](https://circleci.com/gh/garnachod/random_images.svg?style=svg)](https://app.circleci.com/pipelines/github/garnachod/random_images)
```
echo "unsplash=real_unsplash_access_key" > docker/real.env
make docker-build
make docker-run
make docker-test
```


#### local curls api
```
curl localhost:3000/v0/login -H "Authorization: Basic YWRtaW46cGFzc3dvcmQ="
curl "localhost:3000/v0/images/random?x=800&y=600" -H "Authorization: Bearer <jwt_from_login>"
```

#### coverage
```
ok      github.com/garnachod/random_images/internal/user        0.007s  coverage: 95.2% of statements
ok      github.com/garnachod/random_images/internal/image       0.009s  coverage: 100.0% of statements
```