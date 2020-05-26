doubleMe x = x + x

doubleUs x y = x*2 + y*2

doubleSmallNumber x = if x > 100 then x else x*2

conanO'Brien = "It's a-me, Conan O'Brien!" 


hello ::Show a=>  a -> String
-- hello "lzy" = "hello my boy!"
-- hello "123" = "hello 123456"
hello name = "hello "++(show name)