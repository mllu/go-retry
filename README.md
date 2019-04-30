
### Proposed API

- phase 1: default backoff retry with max retries = 5 and with backoff interval `[1, 2, 4, 8, 16]` in second
    ```
    retry.Do(ctx, funcToRetry, ...opts)
    ```

- phase 2: retry with customized BackOff policy
    ```
    retry.Backoff(ctx, funcToRetry, BackOff)
    ```
