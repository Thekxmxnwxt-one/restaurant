CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role_id INTEGER REFERENCES role(id),
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE login_log (
    id SERIAL PRIMARY KEY,
    employee_id INTEGER REFERENCES employee(id),
    login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(50)
);

CREATE TABLE restaurant_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    capacity INTEGER NOT NULL,
    status VARCHAR(20) CHECK (status IN ('available', 'occupied', 'reserved')) DEFAULT 'available'
);

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) UNIQUE,
    email VARCHAR(100) UNIQUE,
    username VARCHAR(50) UNIQUE,
    password_hash VARCHAR(255),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE menu (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    image_url VARCHAR(255),
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category VARCHAR(50),
    available BOOLEAN DEFAULT TRUE
);

CREATE TABLE restaurant_order (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customer(id),
    employee_id INTEGER REFERENCES employee(id), -- สำหรับ dine-in
    table_id INTEGER REFERENCES restaurant_table(id), -- null ได้ถ้า online
    order_type VARCHAR(20) CHECK (order_type IN ('dine-in', 'pickup', 'delivery')) DEFAULT 'dine-in',
    status VARCHAR(20) CHECK (status IN ('open', 'done', 'cancelled')) DEFAULT 'open',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP
);

CREATE TABLE order_item (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES restaurant_order(id) ON DELETE CASCADE,
    menu_id INTEGER REFERENCES menu(id),
    quantity INTEGER NOT NULL,
    status VARCHAR(20) CHECK (status IN ('pending', 'cooking', 'ready', 'served')) DEFAULT 'pending',
    note TEXT
);

CREATE TABLE kitchen_status (
    id SERIAL PRIMARY KEY,
    order_item_id INTEGER REFERENCES order_item(id) ON DELETE CASCADE,
    updated_by INTEGER REFERENCES employee(id),
    status VARCHAR(20) CHECK (status IN ('pending', 'cooking', 'ready', 'served')),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES restaurant_order(id) UNIQUE,
    amount DECIMAL(10,2) NOT NULL,
    payment_type VARCHAR(20) CHECK (payment_type IN ('cash', 'card', 'QR', 'credit_card', 'mobile_banking')),
    paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE delivery_info (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE REFERENCES restaurant_order(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    delivery_status VARCHAR(20) CHECK (delivery_status IN ('pending', 'delivering', 'delivered', 'cancelled')) DEFAULT 'pending',
    shipped_at TIMESTAMP,
    delivered_at TIMESTAMP
);

CREATE TABLE report (
    id SERIAL PRIMARY KEY,
    report_type VARCHAR(20),
    generated_by INTEGER REFERENCES employee(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data JSONB
);

INSERT INTO role (name) VALUES
('admin'),
('cashier'),
('chef'),
('waiter'),
('customer');

INSERT INTO employee (name, username, password_hash, role_id) VALUES
('Admin User', 'admin', '$2a$10$kLKa1Y8ELsrC4jzsvjAko.vlaF/YZK7M7Hv2.tvRe/1hD3pDEKjnO', 1),
('Chef John', 'chef1', '$2a$10$kLKa1Y8ELsrC4jzsvjAko.vlaF/YZK7M7Hv2.tvRe/1hD3pDEKjnO', 3),
('Waiter Ann', 'waiter1', '$2a$10$kLKa1Y8ELsrC4jzsvjAko.vlaF/YZK7M7Hv2.tvRe/1hD3pDEKjnO', 4);

INSERT INTO restaurant_table (name, capacity) VALUES
('Table 1', 4),
('Table 2', 2),
('Table 3', 6),
('Table 4', 4);

INSERT INTO menu (name, image_url, description, price, category, available) VALUES
('ไก่ทอดคาราอาเกะ', 'https://yayoirestaurants.com/productimages/thumbs/2544_Yayoi-Menubook-2024_880-x-800-px_Set-Menu_.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วเหลือง', 160.00, 'เซต', TRUE),
('ไก่ซอสสามรส', 'https://yayoirestaurants.com/productimages/thumbs/3077_Yayoi-Menubook-2024_880-x-800-px_Set-Menu_.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วเหลือง', 180.00, 'เซต', TRUE),
('ไก่เทอริยากิ', 'https://yayoirestaurants.com/productimages/thumbs/7420_Yayoi-Menubook-2024_880-x-800-px_Set-Menu_.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วเหลือง', 190.00, 'เซต', TRUE),
('หมูผัดกิมจิ', 'https://yayoirestaurants.com/productimages/thumbs/7618_Yayoi-Menubook-2024_880-x-800-px_Set-Menu_.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน หอยและหมึก ถั่วเหลือง', 200.00, 'เซต', TRUE),
('เซตกุ้งและไก่ซอสสามรส', 'https://yayoirestaurants.com/productimages/2524_Yayoi-Menubook-2024_880-x-800-px_Set-Menu_.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว สัตว์น้ำมีเปลือก ถั่วเหลือง', 200.00, 'เซต', TRUE),
('โซเม็ง', 'https://yayoirestaurants.com/productimages/thumbs/5071_Yayoi-Menubook-2024_880-x-800-px_Noodle-.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วเหลือง', 175.00, 'ราเมง', TRUE),
('อุด้งหมูสุกี้ยากี้', 'https://yayoirestaurants.com/productimages/thumbs/8231_ND003_Sukiyaki-Udon.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วลิสง ถั่วเหลือง', 185.00, 'ราเมง', TRUE),
('เทมปุระกุ้ง ต้มยำ ราเม็ง', 'https://yayoirestaurants.com/productimages/thumbs/9900_Yayoi-Menubook-2024_880-x-800-px_Noodle---.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว สัตว์น้ำมีเปลือก ถั่วเหลือง', 199.00, 'ราเมง', TRUE),
('ยากิโซบะกระทะร้อน', 'https://yayoirestaurants.com/productimages/1723_05%20-%20880%20x%20800%20Noodle%20-%20%20(Rev.).jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน ถั่วเหลือง', 149.00, 'ราเมง', TRUE),
('โซบะเย็น', 'https://yayoirestaurants.com/productimages/thumbs/3803_Yayoi-Menubook-2024_880-x-800-px_Noodle-.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วเหลือง', 149.00, 'ราเมง', TRUE),
('ชิกกี้ เบนโตะ', 'https://yayoirestaurants.com/productimages/thumbs/3799_3.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วลิสง ถั่วเหลือง', 245.00, 'เบนโตะ', TRUE),
('ซากานะ เบนโตะ', 'https://yayoirestaurants.com/productimages/thumbs/3279_1.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วลิสง ถั่วเหลือง', 255.00, 'เบนโตะ', TRUE),
('ซาชิมิ เดอลุกซ์ เบนโตะ', 'https://yayoirestaurants.com/productimages/thumbs/5104_2_0.jpg', 'เมนูนี้มีส่วนผสมของไข่ ปลา กลูเตน นมวัว ถั่วลิสง ถั่วเหลือง', 385.00, 'เบนโตะ', TRUE),
('ชาสตรอว์เบอร์รี', 'https://yayoirestaurants.com/productimages/7756_Yayoi-Menubook-2024_880-x-800-px_Drink_.jpg', 'เมนูนี้มีส่วนผสมของนมวัว', 60.00, 'เครื่องดื่ม', TRUE),
('มัทฉะ ลาเต้', 'https://yayoirestaurants.com/productimages/4490_Yayoi-Menubook-2024_880-x-800-px_Drink_-.jpg', 'เมนูนี้มีส่วนผสมของนมวัว', 70.00, 'เครื่องดื่ม', TRUE),
('สตรอว์เบอร์รี โซดา', 'https://yayoirestaurants.com/productimages/3425_Yayoi-Menubook-2024_880-x-800-px_Drink_-.jpg', 'เมนูนี้มีส่วนผสมของนมวัว', 60.00, 'เครื่องดื่ม', TRUE),
('วาราบิโมจิ', 'https://yayoirestaurants.com/productimages/8605_%20-%20Copy.jpg', 'เมนูนี้มีส่วนผสมของกลูเตน ถั่วเปลือกแข็ง ถั่วเหลือง', 79.00, 'ของหวาน', TRUE),
('ถั่วแดงร้อนญี่ปุ่นโมจิย่าง', 'https://yayoirestaurants.com/productimages/9330_%20-%20Copy.jpg', 'เมนูนี้มีส่วนผสมของกลูเตน', 65.00, 'ของหวาน', TRUE);

INSERT INTO customer (name, phone, email, username, password_hash, address) VALUES
('สมชาย ลูกค้า', '0891234567', 'somchai@example.com', 'somchai', '$2b$10$Wnq4CkKm99pPWkCDbPtF8eWrKgRL/U9dN7TqLwEto1C1R8R5pUiQC', '99/9 ถนนสุขใจ เขตบางเขน กทม.');

INSERT INTO restaurant_order (customer_id, order_type, status)
VALUES (1, 'delivery', 'open');

INSERT INTO order_item (order_id, menu_id, quantity, note)
VALUES 
(1, 1, 2, 'ไม่เผ็ด'),
(1, 14, 1, 'ใส่น้ำแข็งน้อย');

INSERT INTO delivery_info (order_id, address)
VALUES (1, '99/9 ถนนสุขใจ เขตบางเขน กทม.');

INSERT INTO payment (order_id, amount, payment_type)
VALUES (1, 130.00, 'mobile_banking');
