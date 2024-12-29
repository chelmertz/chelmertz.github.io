---
date: "2011-12-19 23:17"
title: "Properly display MySQL result in console"
permalink: "properly-display-mysql-result-console"
tags: [mysql]
---

If your query results in a result set with many columns, you might end up with totally unreadable jibberish since the console defaults to print the columns horizontally. To display the data vertically instead, with support for really many columns, end the query with \G:

```sql
select * from users \G
```

... instead of

```sql
select * from users; 
```
