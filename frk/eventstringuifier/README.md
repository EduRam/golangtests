Parses this kind of string 

```
abc=[def=ghi][jkl=mno]]
```

into...

```
EventName = abc
  EventAttr = def
  EventValue = ghi
  EventAttr = jkl
  EventValue = mno
```
