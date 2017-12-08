drop table stat;
drop table stat_type;
drop table beer;
drop table vital_statistic;
drop table ingredient;
drop table characteristic_ingredient;
drop table style_tag;
drop table tag;
drop table comparison_style;
drop table style;
drop table category;

create table category (
  category_name varchar(100) primary key,
  category_number varchar(8),
  category_description varchar(1000),
  classis_category_flag varchar(1) -- boolean flag, Y or N
);

create table style (
  style_name varchar(100) primary key,
  style_letter varchar(8),
  category_name varchar(1000) not null,
  overall_impression varchar(1000),
  appearance varchar(1000),
  aroma varchar(1000),
  flavor varchar(1000),
  mouthfeel varchar(1000),
  history varchar(1000),
  comment_text varchar(1000),
  constraint fk_category foreign key (category_name) references category(category_name)
);

create table comparison_style (
  style_name varchar(100) not null,
  use_as_compared_style_name varchar(100) not null,
  style_comparison_note varchar(1000),

  constraint fk_style foreign key (style_name) references style(style_name),
  constraint fk_use_as_comparision_style foreign key (use_as_compared_style_name) references style(style_name),
  constraint comparison_style_pk primary key (style_name, use_as_compared_style_name)
);

drop table tag_type;
create table tag_type (
  tag_type_name varchar(100) primary key
);

create table tag (
  tag_name varchar(100) primary key,
  tag_type varchar(100) not null,
  tag_description varchar(1000),
  constraint fk_tag_type foreign key (tag_type) references tag_type(tag_type_name)
);

create table style_tag (
  tag_name varchar(100) not null,
  style_name varchar(100) not null,
  constraint style_tag_pk primary key (tag_name, style_name)
);

create table characteristic_ingredient (
  style_name varchar(100) not null,
  ingredient varchar(100) not null,
  ci_note varchar(1000),
  constraint fk_style foreign key (style_name) references style(style_name),
  constraint characteristic_ingredient_pk primary key (style_name, ingredient)
);

create table ingredient (
  ingredient_name varchar(100) primary key
);

create table vital_statistic (
  style varchar(100) not null,
  statistic_type varchar(50) not null,
  statistic_lower float,
  statistic_upper float,
  constraint vital_statistics primary key (style, statistic_type)
);

CREATE TABLE beer(
  beer_id INTEGER PRIMARY KEY,
  beer_name VARCHAR(100),
  water_type VARCHAR(100),
  water_amount INTEGER,
  gallons_brewed INTEGER,
  yeast_culture_state VARCHAR(100),
  yeast_culture_starter_flag VARCHAR(100),
  yeast_culture_type VARCHAR(100),
  yeast_culture_brand VARCHAR(100),
  yeast_culture_amount INTEGER,
  yeast_nutrient_type VARCHAR(100),
  yeast_nutrient_amount INTEGER,
  carbonation_form VARCHAR(100),
  carbonation_vol INTEGER,
  priming_sugar_type VARCHAR(100),
  priming_sugar_amount INTEGER,
  boiling_time INTEGER,
  brewing_date DATETIME,
  bottling_date DATETIME
);

CREATE TABLE stat_type(
  stat_type_name VARCHAR(20) primary key,
  measuring_unit VARCHAR(20)
);

CREATE TABLE stat(
  stat_value FLOAT,
  beer_id REFERENCES beer(beer_id),
  stat_type_id REFERENCES stat_type(stat_type_id)
);
