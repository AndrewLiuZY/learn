diet :: (Floating a,Show a)=>a->String
diet a = "You weight is "++show(a)++"\n"
         ++"You should eat protein : "++show(protein )++"g         Calories : "++show(proteinCol )++"kcol\n"
         ++"               Carbohydrate : "++show(carbo )++"g    Calories : "++show(carboCol )++"kcol\n"
         ++"               fat : "++show(fat )++"g Calories : "++show(fatCol )++"kcol\n"
         ++"total calories : "++show(sum [proteinCol,fatCol,carboCol])++"kcol\n"
         where protein  = a * 2
               fat  = a * 0.8
               carbo  = a * 2.5
               proteinCol  = (protein )*4
               fatCol  = (fat )*9
               carboCol = (carbo )*4

putDiet :: (Floating a,Show a)=>a->IO()
putDiet a = putStr (diet a)