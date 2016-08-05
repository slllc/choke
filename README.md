# Package choke
###Creates a choke point where actions being performed by multiple goroutines are serialized


To create a new choke: 

> choke := New(depth)<br>

Depth may need to be > 0 to prevent deadlock under some circumstances.

---
You then need to start the choke's goroutine:

> go choke.Doer(ctx)

You can use context.TODO() if you're not using contexts (which if you aren't, you might want to reconsider)
---
to execute code under the choke, wrap it in a call to Do:

...
>err = choke.Do(func()error{<br>
>    // function body goes here,<br>
>    // what ever is returned from this func<br>
>    // will be returned from choke.Do()<br>
>    return nil<br>
>})<br>

...

