# slack notification example

this is slack notification example.

it's including these features

- test with dummy slack client
- test with mock dispatcher

## Dispatcher?

Dispatcher is the object that dispatching the way of notification.

So, If app running with Accessed() function, looks like this.

```go
func Accessed(ctx context.Context, dispatcher dispatcher.Dispatcher, skipped bool) error {
	fmt.Pprintln("before accessed")
	if skipped {
		return nil
	}

	dispatcher.DispatchAccessed(ctx) // this, calling slack notification
	fmt.Println("after accessed")
	return nil
}
```

For example, the situation is changed in future, the media of notifiation will be not slack. If you are writing code using slack client code directly, then, all dependents code(such like Accessed()), rewriting (or code modification) is needed.

But, if using dispatcher, the work like this is not existed.
