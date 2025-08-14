execution flow for func => at first init func. then main func

argument vs parameter

```
func add(a int, b int) int { // here we are receiving value so this is parameter
return a+b
}

add(10,20) //here we're pass value so this is argument
```

First-order Function: সাধারণ ফাংশন যা অন্য ফাংশনকে আর্গুমেন্ট হিসেবে নেয় না এবং রিটার্নও করে না।

Higher-order Function: যে ফাংশন অন্য ফাংশনকে argument হিসেবে নেয় বা ফাংশন রিটার্ন করে বা দুইটাই করে।

First-class Function / First-class Citizen: ফাংশনকে Go-তে ভ্যালু হিসেবে ধরতে পারা যায় — যেমনঃ ভ্যারিয়েবলে রাখা, আর্গুমেন্ট হিসেবে পাঠানো, রিটার্ন করা যায়।
Go ফাংশন = First-class Citizen

Callback Function: যে ফাংশনকে অন্য ফাংশনের ভিতরে argument হিসেবে দিয়ে, পরে কোনো অবস্থায় কল করা হয়।

Argument: ফাংশন কল করার সময় যেসব মান পাঠানো হয়, সেগুলো argument।

Parameter: ফাংশন ডিক্লার করার সময় যেসব ভ্যারিয়েবল রাখা হয় মান ধরার জন্য, সেগুলো parameter।
