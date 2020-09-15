import Data.List

doubleMe x = x + x

doubleUs x y = x*2 + y*2

doubleSmallNumber x = if x > 100 then x else x*2

conanO'Brien = "It's a-me, Conan O'Brien!" 


hello ::Show a=>  a -> String
-- hello "lzy" = "hello my boy!"
-- hello "123" = "hello 123456"
hello name = "hello "++(show name)

--模式匹配
topThree ::(Show a)=>[a]->String
topThree [] = "null array!"
topThree  (a:[]) = "first is "++show(a)
topThree  (a:b:[]) = "first is "++show(a)++" second is "++show(b)
topThree  (a:b:c:[]) = "first is "++show(a)++" second is "++show(b)++" third is "++show(c)
topThree  (a:b:c:_) = "first is "++show(a)++" second is "++show(b)++" third is "++show(c)

--数组可表示为(a:b:c) or [a,b,c] []中只能匹配单个元素
length'lzy ::(Num a)=>[b]->a
length'lzy []=0
length'lzy (_:x)=1 + length'lzy x

sum'lzy ::(Num a)=>[a]->a
sum'lzy []=0
sum'lzy (a:b) = a+sum'lzy b

-- as模式 @前的名称可代表被分割的值的整体
capital :: String->String
capital [] = "is null string"
capital str@(a:_) = "The first char of "++str++" is "++ [a]

-- bmi ::(RealFloat a) => a => a => a
-- bmi weight height = weight / height ^ 2

bmiTell :: (RealFloat a) => a -> a -> String  
bmiTell weight height  
    | bmi <= skinny = "You're underweight, you emo, you!"  
    | bmi <= normal = "You're supposedly normal. Pffft, I bet you're ugly!"  
    | bmi <= fat = "You're fat! Lose some weight, fatty!"  
    | otherwise = "You're a whale, congratulations!"
    where bmi = weight / (height ^ 2)
          skinny = 18.5  
          normal = 25.0  
          fat = 30.0

max' :: Ord a => a -> a -> a
max' a b
    | a > b = a
    | otherwise = b

sumDoubleTuple :: (Num a)=>(a,a)->(a,a)->a
sumDoubleTuple (a,b) (c,d)=
    let fstSum = a+b
        sndSum = c+d
    in fstSum+sndSum

maximum' :: (Ord a) => [a] -> a  
maximum' [] = error "maximum of empty list"  
maximum' [x] = x 
maximum' (x:xs)   
    | x > maxTail = x  
    | otherwise = maxTail  
    where maxTail = maximum' xs

maximum'1 :: (Ord a) => [a] -> a  
maximum'1 arr = case arr of [] ->  error "maximum of empty list"  
                            [x] -> x
                            (x:xs) -> max x (maximum'1 xs)

replicate' :: (Num a , Ord a )=> a->b->[b]
replicate' num ele
    | num < 0 = error "num must more than 0" 
    | num == 0 = []
    | otherwise = ele:replicate' (num-1) ele

take' :: (Num a , Ord a )=> a->[b]->[b]
take' num _
    | num<=0 = []
take' _ [] = []
take' num (x:xs) = x:take' (num-1) xs

reverse' ::[a]->[a]
reverse' [] = []
reverse' (x:xs) = reverse' xs ++ [x]

repeat' :: a->[a]
repeat' a = a:repeat' a

zip' ::[a]->[b]->[(a,b)]
zip' _ [] = []
zip' [] _ = []
zip' (x:xs) (y:ys) = (x,y):zip' xs ys

elem' :: (Eq a)=>a->[a]->Bool
elem' _ [] = False
elem' a (x:xs)
    | a==x = True
    | otherwise = elem' a xs

quickSort :: (Ord a)=>[a]->[a]
quickSort [] = []
quickSort (x:xs) = smallerSorted ++ [x] ++ biggerSorted
    where smallerSorted = quickSort [a|a<-xs,a<=x]
          biggerSorted = quickSort [a|a<-xs,a>x]

applyTwice :: (a->a)->a->a
applyTwice f x = f (f x)

zipWith' :: (a->b->c)->[a]->[b]->[c]
zipWith' _ _ [] = []
zipWith' _ [] _ = []
zipWith' f (x:xs) (y:ys) = f x y : zipWith' f xs ys

flip' :: (a->b->c)->b->a->c
flip' f x y = f y x 

map' ::(a->b)->[a]->[b]
map' _ [] = []
map' f (x:xs) = f x : map' f xs

filter' :: (a->Bool)->[a]->[a]
filter' _ [] = []
filter' f (x:xs)
    | f x = x:filter' f xs
    | otherwise = filter' f xs

largestDivisible :: (Integral a) => a  
largestDivisible = head (filter p [100000,99999..])  
    where p x = x `mod` 3829 == 0

chain ::(Integral a)=>a->[a]
chain 1 =[1]
chain n
    | even n = n:chain (div n 2)
    | odd n = n:chain (n*3+1) 

elem'foldl ::(Eq a)=>a->[a]->Bool
elem'foldl a xs = foldl (\acc x->if a==x then True else acc) False xs

map'foldr ::(a->b)->[a]->[b]
map'foldr f xs = foldr (\x acc->(f x):acc) [] xs

map'foldl ::(a->b)->[a]->[b]
map'foldl f xs = foldl (\acc x->(f x):acc) [] xs

search :: (Eq a) => [a] -> [a] -> Bool
search needle haystack =
    let nlen = length needle
    in foldl (\acc x -> if take nlen x == needle then True else acc) False (tails haystack)