(func
	(local i32 i32)
	(local.get 0)
	(i32.const 1)
	(i32.add)
	(local.set 0)
	(local.get 0)
	(local.set 1)
	(local.get 0)
	(i32.const 1)
	(i32.add)
	(local.set 1)
	(local.get 0)
	(local.get 1)
	(i32.add)
	(drop)
)

(start 0)

