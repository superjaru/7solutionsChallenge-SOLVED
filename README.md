## 1. จงหาเส้นทางที่มีค่ามากที่สุด


- ซึ่งจะอยู่ใน Array ดังต่อไปนี้ `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]`
- โดยเส้นทางที่มีค่ามากที่สุดจะเป็นตามจุดสีแดง
- แต่ละ node ห้ามย้อนกลับ (ต้องขึ้นลงเป็นทางเดียว) และเชื่อมกัน
- คำตอบให้อยู่ในรูปของ จำนวนรวมของเส้นทางที่ผ่าน ซึ่งจากตัวอย่างคือ `237`

ให้เขียนโปรแกรมภาษา GO โดยใช้ input จากไฟล์นี้ /hard.json และแสดงผลเป็นค่าที่ได้จากการคำนวณ

# Test case

- input = `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]` output = `237`
- input = /hard.json output = `7273`

## 2. จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ

ให้ถอดรหัสตัวอักษรตามตัวอย่างด้านล่าง

```
สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา

ตัวอย่างเช่น ตัวเลขชุด “410233“ จะเข้ารหัสได้เป็น “LLRR=“

อธิบายเพิ่มเติม จากตัวเลขชุด “410233“

เลขคู่แรกคือ 4 และ 1      => ตัวเลขซ้าย (4) มากกว่า ตัวเลขขวา (1)   => “L“
เลขคู่ถัดมาคือ 1 และ 0   => ตัวเลขซ้าย (1) มากกว่า ตัวเลขขวา (0)   => “L“
เลขคู่ถัดมาคือ 0 และ 2   => ตัวเลขซ้าย (0) น้อยกว่า ตัวเลขขวา (2)  => “R“
เลขคู่ถัดมาคือ 2 และ 3   => ตัวเลขซ้าย (2) น้อยกว่า ตัวเลขขวา (3)  => “R“
เลขคู่ถัดมาคือ 3 และ 3   => ตัวเลขซ้าย (3) เท่ากับ ตัวเลขขวา (3)  => “=“

“LLRR=” สามารถแปลงได้เป็น “410233” หรือ “210122“ ก็ได้ 
ผลรวมตัวเลขทุกตัวของ 410233 = 4 + 1 + 0 + 2+. 3 + 3 = 13
ผลรวมตัวเลขทุกตัวของ 210122 = 2 + 1 + 0 + 1 + 2 + 2  = 8
คำตอบคือ เลขชุด 210122 เนื่องจาก ผลรวมของทุกตัวเลขมีค่าน้อยที่สุด
```

ให้เขียนโปรแกรม ภาษา GO เพื่อรับข้อมูลข้อความที่เข้ารหัสแล้ว จาก keyboard (encoded) และให้แปลงกับเป็นตัวเลขชุด ที่มีผลรวมของทุกตัวเลข มีค่าน้อยที่สุดแสดงผลเป็น ตัวเลขชุด ชุดนั้น

# Test case

- input = `LLRR=` output = `210122`
- input = `==RLL` output = `000210`
- input = `=LLRR` output = `221012`
- input = `RRL=R` output = `012001`

## 3. พาย ไฟ ได - Pie Fire Dire  

โจทย์คือ ให้รายชื่อของเนื้อหลายชนิด ปะปนกันเช่น

```Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.```

`ทุกคำเป็นชื่อชนิดเนื้อทั้งหมด` โดยที่ไม่ต้องสนใจ , . หรือ space

จงหารายชื่อเนื้อทั้งหมด และระบุจำนวนของเนื้อแต่ละชนิด

ใช้ข้อความจาก <https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text>

- Your project must use Go, Go module, and HTTP framework (GRPC is plus)
- You create a JSON API at this endpoint /beef/summary
- The JSON API must count the number of beef
- We encourage you to write tests, which we will give you some extra score
- We will give you an extra score if you focus on performance.

--- sample response --

```json
{
    "beef": {
        "t-bone": 4,
        "fatback": 1,
        "pastrami": 1,
        "pork": 1,
        "meatloaf": 1,
        "jowl": 1,
        "enim": 1,
        "bresaola": 1
    }
}
```

สรุป problem solving ของทั้ง 3 ข้อ เพื่อให้ง่าย ต่อการอ่านcodeครับ


Q1 :  Solve โดยการให้ read file จาก hard.json และทำการ calculate sum ของแต่ละ node ไล่ลงมาตั้งแต่ level 1 จนถึงล่างสุด แล้วหาผลรวมที่มากที่สุดจาก node ที่  bottom level ครับ

Q2 : ต้องบอกว่าข้อนี้แอบใช้เวลานานมากครับ สุดท้ายแล้วเลย solve ด้วยการ hardcoding สร้าง encoded ขึ้นมาใหม่แล้วเทียบหาตัวที่ผลรวมน้อยที่สุด (ถ้า input มากกว่า 7 ตัวจะใช้เวลานานขึ้น) และเพิ่มเติมคือ input จาก user มีการเช็ค valid ให้ใส่ได้แค่ format R , L , = เท่านั้น 


Q3 : ข้อนี้ทำตรงๆครับ ใช้เป็น http framework ซึ่งมีการทำเพิ่มเติมในส่วน extra scores เพิ่มครับ คือ
   
    - We encourage you to write tests, which we will give you some extra score
       : มีเขียนเทสเพิ่มครับในไฟล์ main_test.go ซึ่งมีการเทส server , request , function ต่างๆ

    - We will give you an extra score if you focus on performance.
       : ใช้ lib strings ซึ่งคือ strings.FieldsFunc ในการ filter พวก ,. ซึ่งทำให้ performance ดีขึ้นครับ เพราะไม่ได้ loop iterate ทีละ character 
