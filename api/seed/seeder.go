package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/rezyfr/go-blog/api/models"
	"log"
)

var user = models.User{
	Name:     "Fidriyanto Rizkillah",
	Email:    "frotylatz@gmail.com",
	Password: "p@ssw0rd",
}

var post = models.Post{
	Guid:        "https://medium.com/p/f1da2427db8c",
	Link:        "https://medium.com/cicil-tech/a-pro-fun-sional-internship-f1da2427db8c?source=rss-938df2b89d90------2",
	Title:       "A Pro-fun-sional Internship",
	Content:     "\n<p>Hello! My name is Rezy, a final-year Computer Science student from Universitas Padjadjaran who has a high-interest in android development. On this summer holiday, I was intended to do an internship to get more work experience and extra allowance :) and luckily Cicil was opening an internship batch at that time.</p>\n<p>So from June to September 2019, I have given a chance to do an internship at Cicil.co.id. Now, I’m gonna tell you how this internship has gone beyond my expectations since the first day I walk into a room so-called ‘office’.</p>\n<figure><img alt=\"\" src=\"https://cdn-images-1.medium.com/max/1024/1*x33olsSBh5LIJseSuf-W0A.jpeg\"></figure><p>Cicil accepted me as an android developer intern around 3 weeks before the internship starts after submitting the application letter and doing an online interview.</p>\n<p>On the first day, I was escorted to some kind of meeting room and gathered with another intern which surprisingly came from several top universities. This makes me kinda nervous since I’m not coming from a top university. We introduced each other and that was our first bonding, everybody was so nice. We talk like we have already known each other for a while.</p>\n<p>After that, HR introduces us to everything we need to know about Cicil, about how the company starts, the dos-don’ts, the culture and much else. After the introduction finished, all the interns are handed over to each mentor and we get to know each other a little bit. We finished that day by playing a board game with our CTO. The impression I get on my first day is, “How can working be this fun?”.</p>\n<p>In the first 2 weeks, it was more like an adjustment for me, I was taught how the android development works, like what library they use, what language, what kind of architecture, etc. Here, they use Java with MVP architecture as their development framework but they also prepared to migrate to MVVM + Kotlin. I learned everything I need to have for becoming a professional android developer. From basics to intermediate things. And after the on-boarding finished, I was assigned to the authentication part of the ambassador application.</p>\n<p>It is not that difficult to implement the mockup to the application and make it works. The most challenging part is to maintain and optimize my code quality since I have never used any architecture since the first time I develop an android application and at that time I realized how ‘ugly’ my code is. Gladly my mentor was such a great teacher, he always reviews my code and really teaches me to always optimize my code. After I finished the ambassador project, my mentor taught me to use every library that commonly used by every application nowadays and implement it on the ambassador project.</p>\n<p>I finished all the mentor’s requirements for the ambassador project in 1,5 months. After that, I assigned to research how to integrate LinkedIn API to our application and I wrote an API Documentation that going to be used by the company.</p>\n<p>After that research, it was kinda a lot of doing nothing and waiting for a new assignment. I used this spare time to learn Kotlin and also doing my college works sometimes. I also go to Cohive from time to time to take a nap or play billiard, yes our office is next to a coworking space that has a lot of cool facilities.</p>\n<p>I was really enjoying this 3 months working on such a fun company. I learned a lot, like REALLY A LOT. But not just that, I’m also having fun every day I go to work. There’s so much thing that makes the employee love working there, like afternoon snacks, monthly treat, flexible work hours and many more. We intern also having a really great time, we get to lunch with the CEO, play a VR game as a farewell event and we also given a lot of Cicil merchandise.</p>\n<p>I would really recommend every college student to apply for an internship here in Cicil every time they have a batch intern opening. I can guarantee that you will enjoy every moment you have there. For me, 3 months is really not enough. There’s still so much I can learn from everyone at Cicil.co.id, be it technical or no.</p>\n<p>After I finish my study I would love it if Cicil could be the first stone I step on my professional career.</p>\n<img src=\"https://medium.com/_/stat?event=post.clientViewed&amp;referrerSource=full_rss&amp;postId=f1da2427db8c\" width=\"1\" height=\"1\" alt=\"\"><hr>\n<p><a href=\"https://medium.com/cicil-tech/a-pro-fun-sional-internship-f1da2427db8c\">A Pro-fun-sional Internship</a> was originally published in <a href=\"https://medium.com/cicil-tech\">CICIL Tech</a> on Medium, where people are continuing the conversation by highlighting and responding to this story.</p>\n",
	Description: "\"\\n<p>Hello! My name is Rezy, a final-year Computer Science student from Universitas Padjadjaran who has a high-interest in android development. On this summer holiday, I was intended to do an internship to get more work experience and extra allowance :) and luckily Cicil was opening an internship batch at that time.</p>\\n<p>So from June to September 2019, I have given a chance to do an internship at Cicil.co.id. Now, I’m gonna tell you how this internship has gone beyond my expectations since the first day I walk into a room so-called ‘office’.</p>\\n<figure><img alt=\\\"\\\" src=\\\"https://cdn-images-1.medium.com/max/1024/1*x33olsSBh5LIJseSuf-W0A.jpeg\\\"></figure><p>Cicil accepted me as an android developer intern around 3 weeks before the internship starts after submitting the application letter and doing an online interview.</p>\\n<p>On the first day, I was escorted to some kind of meeting room and gathered with another intern which surprisingly came from several top universities. This makes me kinda nervous since I’m not coming from a top university. We introduced each other and that was our first bonding, everybody was so nice. We talk like we have already known each other for a while.</p>\\n<p>After that, HR introduces us to everything we need to know about Cicil, about how the company starts, the dos-don’ts, the culture and much else. After the introduction finished, all the interns are handed over to each mentor and we get to know each other a little bit. We finished that day by playing a board game with our CTO. The impression I get on my first day is, “How can working be this fun?”.</p>\\n<p>In the first 2 weeks, it was more like an adjustment for me, I was taught how the android development works, like what library they use, what language, what kind of architecture, etc. Here, they use Java with MVP architecture as their development framework but they also prepared to migrate to MVVM + Kotlin. I learned everything I need to have for becoming a professional android developer. From basics to intermediate things. And after the on-boarding finished, I was assigned to the authentication part of the ambassador application.</p>\\n<p>It is not that difficult to implement the mockup to the application and make it works. The most challenging part is to maintain and optimize my code quality since I have never used any architecture since the first time I develop an android application and at that time I realized how ‘ugly’ my code is. Gladly my mentor was such a great teacher, he always reviews my code and really teaches me to always optimize my code. After I finished the ambassador project, my mentor taught me to use every library that commonly used by every application nowadays and implement it on the ambassador project.</p>\\n<p>I finished all the mentor’s requirements for the ambassador project in 1,5 months. After that, I assigned to research how to integrate LinkedIn API to our application and I wrote an API Documentation that going to be used by the company.</p>\\n<p>After that research, it was kinda a lot of doing nothing and waiting for a new assignment. I used this spare time to learn Kotlin and also doing my college works sometimes. I also go to Cohive from time to time to take a nap or play billiard, yes our office is next to a coworking space that has a lot of cool facilities.</p>\\n<p>I was really enjoying this 3 months working on such a fun company. I learned a lot, like REALLY A LOT. But not just that, I’m also having fun every day I go to work. There’s so much thing that makes the employee love working there, like afternoon snacks, monthly treat, flexible work hours and many more. We intern also having a really great time, we get to lunch with the CEO, play a VR game as a farewell event and we also given a lot of Cicil merchandise.</p>\\n<p>I would really recommend every college student to apply for an internship here in Cicil every time they have a batch intern opening. I can guarantee that you will enjoy every moment you have there. For me, 3 months is really not enough. There’s still so much I can learn from everyone at Cicil.co.id, be it technical or no.</p>\\n<p>After I finish my study I would love it if Cicil could be the first stone I step on my professional career.</p>\\n<img src=\\\"https://medium.com/_/stat?event=post.clientViewed&amp;referrerSource=full_rss&amp;postId=f1da2427db8c\\\" width=\\\"1\\\" height=\\\"1\\\" alt=\\\"\\\"><hr>\\n<p><a href=\\\"https://medium.com/cicil-tech/a-pro-fun-sional-internship-f1da2427db8c\\\">A Pro-fun-sional Internship</a> was originally published in <a href=\\\"https://medium.com/cicil-tech\\\">CICIL Tech</a> on Medium, where people are continuing the conversation by highlighting and responding to this story.</p>\\n\"",
	CategoryId:  "{1,2}",
	Thumbnail:   "https://cdn-images-1.medium.com/max/1024/1*x33olsSBh5LIJseSuf-W0A.jpeg",
}

var category = models.Category{
	Categories: "Android Development",
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Category{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).
		AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	post.AuthorID = user.ID

	err = db.Debug().Model(&models.Post{}).Create(&post).Error
	if err != nil {
		log.Fatalf("cannot seed posts table: %v", err)
	}

	err = db.Debug().Model(&models.Category{}).Create(&category).Error
	if err != nil {
		log.Fatalf("cannot seed category table: %v", err)
	}
}
