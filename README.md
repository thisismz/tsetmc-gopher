
# tsetmc-gopher

<div dir="rtl">

# خوش آمدین tsetmc gopher به پروژه متن باز   

<p align="center">
  <img src="https://github.com/thisismz/tsetmc-gopher/blob/main/BLUE_GOPHER.png" width="250" hight="300" title="The GOPHER">
</p>

<div dir="rtl">
این پروژه جهت دریافت لحظه ای  داده ها از سایت سازمان بورس به نشانی tsetmc طراحی شده است هدف نهایی از این پروژه ایجاد بستری مناسب جهت انالیز لحظه ای بازار سرمایه می باشد.

# محتویات
- [قابلیت ها](#قابلیت-ها)
- [نصب](#نصب)
- [نحوه استفاده](#نحوه-استفاده)
- [نمونه خروجی](#نمونه-خروجی)

## قابلیت ها
-  دریافت اطلاعات لحظه ای هر سهم و قابلیت ذخیره سازی بر روی انواع دیتابیس ها
-  قابلیت گرفتن اطلاعات یک سهام مانند اطلاعات معاملات حقیقی و حقوقی واطلاعات معاملات تابلوحجم معاملات
- دریافت اطلاعات فاندامنتال یک نماد شامل(eps, p/e ,حجم مبنا)
- دریافت لحظه لیست نماد های موجود
## نصب
1-اعمال تنظیمات دیتابیس در فایل (config.yaml)
**نکته : تعداد رشته ها جهت دریافت داده از سایت نیز در این فایل میتوانید مشخص کنید**

> system:  
  db-type: 'postgres'  // نوع دیتابیس
  thread: 12 // تعداد رشته

2- اجرای دستور  
> go build

3- اجرای نرم فزار توسط کلیک بر روی فایل خروجی

4-  درصورت اعمال صحیح تنظیمات مربوط به کانکشن دیتابیس با پیام
`Successfully connection`
مواجه خواهید شد

## نحوه استفاده
نرم افزار را اجرا کنید بعد از مواجهه با پیام `enter sym name`نام سهم مورد نظر را در cmd تایپ کنید
## نمونه خروجی 
![نمونه خروجی](https://i.imgur.com/JWfCFAb.png)
## نکته نرم افزار در مرحله الفا می باشد
همانطور که مستحضر هستین نرم افزار کامل نیست و نیاز به افزون قابلیت های متعدد می باشد
چون داده ها به صورت لحظه ای دریافت و در دیتابیس ذخیره می شوند شما می توانید با دستور sql به هر شکل که میخوایید از داده ها استفاده کنید
## TODO
- [ ] دریافت eps
- [ ] دریافت گروه زیر مجموعه سهم
- [ ] ایجاد api
- [ ] ایجاد ui

<div dir="rtl">
