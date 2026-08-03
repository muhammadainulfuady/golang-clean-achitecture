create table
  product (
    id_product int primary key auto_increment,
    name varchar(255) not null,
    stok int not null,
    price decimal(10, 2) not null
  ) ENGINE = InnoDB;