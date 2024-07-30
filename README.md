# OpenFeature Experiments

Experiments related to Configcat and OpenFeature; mostly running through the docs.

## Contents
* `cc-server` - a simple HTTP server that uses the ConfigCat SDK directly.
* `os-server` - a simple HTTP server that uses the OpenFeature SDK.

## ConfigCat Notes

* A `Setting` can be of various types, like text, int, float or bool.
  * A `feature flag` is a setting of type bool.
* A `Config` is a collection of `Settings`.
  * A Config is like an online version of a traditional config file.
* Each environment-config pair has its own SDK Key which must be used to initialize the ConfigCat SDK within your application.

* Can use webhooks to only download `config.json` when there is a change to a flag, see:
  * https://configcat.com/docs/requests/#use-webhooks
  * Lazy polling can also give you fewer `config.json` downloads.

* Evaluate flags in the backend rather than the frontend to avoid excessive `config.json` downloads:
  * https://configcat.com/docs/requests/#call-your-backend-instead-of-the-configcat-cdn

* Use segments to create reusable groups of users for targetting rules.

* Can separate configs (eg frontend flags/backend flags) to make smaller `config.json` downloads
  * https://configcat.com/docs/network-traffic/#separate-your-feature-flags-into-multiple-configs

* The `defaultValue` parameter to the `Get*Value` functions is returned in case of an error.

## ConfigCat Go SDK notes

* `NewCustomClient` allows setting options on the client. Interesting options:
  * `BaseUrl` - if we use a proxy
  * `Logger` - for fiddling with logging
  * `LogLevel` - for filtering log levels
  * `PollInterval` - essential
  * `FlagOverides` - for local testing?
  * `Hooks` - for metrics?

* Can get details about flag evaluation using `Get[TYPE]ValueDetails()` functions:
  * https://configcat.com/docs/sdk-reference/go/#anatomy-of-gettypevaluedetails

* The `OnFlagEvaluated` hook could be used for metrics about flag evaluations.
  * https://configcat.com/docs/sdk-reference/go/#hooks

* Flag overrides look essential for local development:
  * https://configcat.com/docs/sdk-reference/go/#flag-overrides
  * Changing local flags json needs a restart of the app.
  * Can also do it inline when creating the client rather than needing a file on disk.

## OpenFeature Notes

* Evaluation contexts can be set at several levels, can have some values that are always present by setting them on the client
  * https://openfeature.dev/docs/reference/concepts/evaluation-context
