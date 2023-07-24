CREATE TABLE users_p (
    id serial not null unique,
    first_last varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    status varchar(8) not null default 'activo',
    password_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'), 
    uri_img text, 
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE empresas (
    id serial not null unique,
    nombre_empresa varchar(255) not null unique,
    nuemro_dni varchar(15) not null unique,
    description text,
    telefono_oficina varchar(50),
    telefono_celular varchar(50),
    uri_img text,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    delete_at TIMESTAMP
);

CREATE TABLE categories (
    id serial not null unique,
    name_category varchar(255) not null unique,
    uri_img text,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP
);

CREATE TABLE products (
    id serial not null unique,
    id_uuid uuid DEFAULT uuid_generate_v4 () not null unique,
    name_product varchar(255) not null unique,
    description text,
    uri_img text,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP,
    fk_id_category int references categories (id) on delete cascade not null
);

CREATE TABLE unidad_medidas (
    id serial not null unique,
    name_unidad varchar(255) not null unique,
    valor_unidad int,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP
);


CREATE TABLE almacenes (
    id serial not null unique,
    id_uuid uuid DEFAULT uuid_generate_v4 () not null unique,
    nombre_stop varchar(255) not null unique,
    abrevia varchar(4) not null unique,
    ubication text,
    phone varchar(255) not null,
    encargado_principal varchar(255) not null,
    phone_encargado varchar(255) not null,
    uri_img text,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    delete_at TIMESTAMP
);

CREATE TABLE roles (
    id serial not null unique,
    name_role varchar(10) not null unique,
    nivel_role int default 3,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    delete_at TIMESTAMP
);

CREATE TABLE roles_users (
    id serial not null unique,
    nivel_opcion int,
    fk_id_user int references users (id) on delete cascade not null,
    fk_menu_opcion_detalle int references menu_opcions_detalles (id) on delete cascade not null,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    delete_at TIMESTAMP
);

CREATE TABLE users (
    id serial not null unique,
    first_last varchar(255) not null,
    username varchar(255) not null unique,
    telefono varchar(100) not null,
    email varchar(255) not null unique,
    password_hash varchar(255) not null,
    status varchar(8) not null default 'activo',
    uri_img text,
    fk_instituto int references institutos (id) on delete cascade not null,
    fk_rol int references roles (id) on delete cascade not null
);



CREATE TABLE menu_opciones (
  id serial not null unique,
  name_menu varchar(50) not null unique,
  status varchar(8) NOT NULL DEFAULT 'activo'::character varying,
  created_at timestamp(6) DEFAULT now(),
  updated_at timestamp(6) DEFAULT now(),
  delete_at timestamp(6)
);

CREATE TABLE menu_opcions_detalles (
  id serial not null unique,
  name_opcion varchar(50) not null unique,
  description varchar(250) not null,
  icon varchar(50) not null,
  componente_uri varchar(50) not null,
  page_url varchar(200) not null,
  orderby int, 
  type_opcion varchar(1) not null DEFAULT 'L'::character varying, 
  fk_menu_opcion int references menu_opciones (id) on delete cascade not null,
  status varchar(8) NOT NULL DEFAULT 'activo'::character varying,
  created_at timestamp(6) DEFAULT now(),
  updated_at timestamp(6) DEFAULT now(),
  delete_at timestamp(6)
);

CREATE TABLE municipios (
    id serial not null unique,
    codigo_municipio varchar(2) not null unique,
    nombre_municipio varchar(255) not null,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP
);
CREATE TABLE parroquias (
    id serial not null unique,
    codigo_parroquia varchar(3) not null unique,
    nombre_parroquia varchar(255) not null,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP,
    fk_codigo_municipio varchar(2) references municipios (codigo_municipio) on delete cascade not null
);
CREATE TABLE beneficiarios (
    id serial not null unique,
    cedula varchar(12) not null unique,
    nombre_benefi varchar(255) not null,
    fecha_nacimiento date,
    edad int,
    telefono varchar(50),
    correo varchar(100),
    direccion text,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP,
    fk_codigo_municipio varchar(2) references municipios (codigo_municipio) on delete cascade not null,
    fk_codigo_parroquia varchar(3) references parroquias (codigo_parroquia) on delete cascade not null,
    fk_user int references users (id) on delete cascade not null
);



CREATE TABLE inventarios (
    id serial not null unique,
    price float,
    existence int,
    stop_min int,
    stop_max int,
    discount varchar(1) not null default 'S',
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP,
    fk_id_instituto int references institutos (id) on delete cascade not null,
    fk_id_product int references products (id) on delete cascade not null,
    fk_id_unidad_medida int references unidad_medidas (id) on delete cascade not null
);

CREATE TABLE solicitud_tipos (
    id serial not null unique,
    name_solicitud varchar(255) not null unique,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP
);

CREATE TABLE donacione_tipos (
    id serial not null unique,
    nombre_tipo varchar(255) not null unique,
    tipo varchar(20) not null,
    descripcion text,
    fk_id_solicitud_tipo int references solicitud_tipos (id) on delete cascade not null,
    status varchar(8) not null default 'activo',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP
);

CREATE TABLE donaciones (
    id serial not null unique,
    numero_solicitud varchar(15) not null unique,
    fecha_solicitud date not null,
    nombre_centro varchar(255),
    medico_tratante varchar(255),
    observacion text,
    enlace_politico varchar(255),
    remitido varchar(255),
    responsable varchar(255),
    fecha_entrega date,
    status varchar(10) not null default 'pendiente',
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    deleted_at TIMESTAMP,
    fk_benefici varchar(12) references beneficiarios (cedula) on delete cascade not null,
    fk_instituto int references institutos (id) on delete cascade not null,
    fk_donacione_tipo int references donacione_tipos (id) on delete cascade not null,
    fk_user int references users (id) on delete cascade not null
);
CREATE TABLE donacione_detalle (
    id serial not null unique,
    cantidad int,
    descripcion varchar(255),
    fk_numero_solicitud varchar(15) references donaciones (numero_solicitud) on delete cascade not null,
    fk_id_inventario int references inventarios (id) on delete cascade not null,
    created_at TIMESTAMP default now(),
    updated_at TIMESTAMP default now(),
    delete_at TIMESTAMP
);