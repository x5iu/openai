# Yet Another OpenAI Golang SDK

**It's still unstable and likely to undergo breaking changes.**

This is the SDK I personally use to call OpenAI(also Azure) services, and I generate interface invocation code using [defc](https://github.com/x5iu/defc). Particularly, in the implementation of streaming transmission, I utilize the features of Golang channels, allowing for receiving data chunks in a loop using the `for range` construct.