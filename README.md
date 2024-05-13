สรุป problem solving ของทั้ง 3 ข้อ เพื่อให้ง่าย ต่อการอ่านcodeครับ


Q1 :  Solve โดยการให้ read file จาก hard.json และทำการ calculate sum ของแต่ละ node ไล่ลงมาตั้งแต่ level 1 จนถึงล่างสุด แล้วหาผลรวมที่มากที่สุดจาก node ที่  bottom level ครับ

Q2 : ต้องบอกว่าข้อนี้แอบใช้เวลานานมากครับ สุดท้ายแล้วเลย solve ด้วยการ hardcoding สร้าง encoded ขึ้นมาใหม่แล้วเทียบหาตัวที่ผลรวมน้อยที่สุด (ถ้า input มากกว่า 7 ตัวจะใช้เวลานานขึ้น) และเพิ่มเติมคือ input จาก user มีการเช็ค valid ให้ใส่ได้แค่ format R , L , = เท่านั้น 


Q3 : ข้อนี้ทำตรงๆครับ ใช้เป็น http framework ซึ่งมีการทำเพิ่มเติมในส่วน extra scores เพิ่มครับ คือ
   
    - We encourage you to write tests, which we will give you some extra score
       : มีเขียนเทสเพิ่มครับในไฟล์ main_test.go ซึ่งมีการเทส server , request , function ต่างๆ

    - We will give you an extra score if you focus on performance.
       : ใช้ lib strings ซึ่งคือ strings.FieldsFunc ในการ filter พวก ,. ซึ่งทำให้ performance ดีขึ้นครับ เพราะไม่ได้ loop iterate ทีละ character 
