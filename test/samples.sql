select id, `sort` from authors where substr(sort,1,instr(sort,' ')-1) like 'А_';
SELECT distinct substr(sort,1,1) as s FROM authors where substr(sort,1,1) between 'А' and 'Я';
SELECT substr(sort,1,2) as s, count(*) as c FROM authors where sort like 'А%' group by s;
select max(r.c) as m from (SELECT substr(sort,1,2) as s, count(*) as c FROM authors where sort like 'А%' group by s) as r;
select sum(r.c) as sm from (SELECT substr(sort,1,2) as s, count(*) as c FROM authors where sort like 'А%' group by s) as r;
SELECT substr(sort,1,1) as s, count(*) as c FROM authors where sort like 'ГОЛБ%' and (substr(sort,1,1) between 'А' and 'Я') group by s;
SELECT substr(sort,1,1) as s, count(*) as c FROM authors as a where sort like '%' and ((substr(sort,1,1) between 'А' and 'Я') or (substr(sort,1,1) between 'A' and 'Z')) group by s;
