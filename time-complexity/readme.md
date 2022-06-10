# Runtime Complexity


## For Loop

   a. **Loop dalam loop**

  contoh:
  ```go
  for i:=0; i<5; i++{
    for j:=0; j<4; j++{
      //do here...
    }
  }
  ```

   Time Complexity : `(O(n*m))`

   - n = panjang array / data
   - m = loop dalam loop
  
   b. **Loop + Loop**

contoh:
   ```go
    func a(){
        for i:=0; i<5; i++{
            //do here...
        }

        for j:=0; j<5; j++{
            //do here...
        }
    }
   ```



   Time Complexity : `O(l1 + l2 + ... + ln)`
   