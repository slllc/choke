
# Package choke
### Creates a choke point where actions being performed by multiple goroutines are serialized according to width.

``` Note that this behavior has changed from the original version```


To create a new choke:

> choke := New(width)<br>

Width sets the degree of parallelism. Width = 1 makes a choke a mutex.

---
You then need to start the choke's goroutine(s):

> choke.Start(ctx)

Just for fun, you can combine the New and Start:

choke := New(width).Start(ctx)

##### You can use context.TODO() if you're not using contexts (which if you aren't, you might want to reconsider)
---
### to execute code under the choke, wrap it in a call to Do:

...
>err = choke.Do(func()error{<br>
>    // function body goes here,<br>
>    // what ever is returned from this func<br>
>    // will be returned from choke.Do()<br>
>    return nil<br>
>})<br>

...

