## Context Values

Just use that if it is a request scope. If it is something not related to that request, you should not use.

Everything we put in context is invisible, and it is easy to full context with a lot of values. If we put to much thing there will be hard to follow what the code is doing. 

**Rules to use context values**: 

1. Whatever is in the context should be request specific
2.  Whatever is in the context should be information that is extra info that is usefull, but thats not impact the flow of the program.