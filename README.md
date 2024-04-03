# Yet Another OpenAI Golang SDK

**It's still unstable and likely to undergo breaking changes.**

This is the SDK I personally use to call OpenAI(also Azure) services, and I generate interface invocation code using [defc](https://github.com/x5iu/defc). Particularly, in the implementation of streaming transmission, I utilize the features of Golang channels, allowing for receiving data chunks in a loop using the `for range` construct.

## Update Model Constants

Use the following command to update available models:

```shell
OPENAI_BASE_URL=$YOUR_BASE_URL OPENAI_API_KEY=$YOUR_API_KEY go generate -run=generate_models_file
```

It will call the `/models` endpoint to query the currently available models and update them as constants in the `constants.gen.go` file for use.