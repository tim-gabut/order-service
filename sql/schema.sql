-- ==========================================
-- ORDERS (HEADER)
-- ==========================================
CREATE TABLE orders
(
    id            BIGSERIAL PRIMARY KEY,
    code          VARCHAR(30)  NOT NULL UNIQUE,
    customer_name VARCHAR(100) NOT NULL,
    order_type    VARCHAR(20)  NOT NULL, -- DINE_IN / TAKEAWAY
    table_number  VARCHAR(10),           -- nullable utk TAKEAWAY
    pickup_time   TIMESTAMP,             -- optional
    status        VARCHAR(20)  NOT NULL, -- NEW / ACCEPTED / READY / COMPLETED
    total_amount  BIGINT       NOT NULL DEFAULT 0,
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- Index buat sering query status di dashboard barista
CREATE INDEX idx_orders_status ON orders (status);

-- ==========================================
-- ORDER ITEMS (DETAIL)
-- ==========================================
CREATE TABLE order_items
(
    id           BIGSERIAL PRIMARY KEY,
    order_id     BIGINT       NOT NULL,
    product_id   BIGINT       NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    quantity     INT          NOT NULL,
    price        BIGINT       NOT NULL,
    subtotal     BIGINT       NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order
        FOREIGN KEY (order_id)
            REFERENCES orders (id)
            ON DELETE CASCADE
);

-- Index untuk lookup cepat by order
CREATE INDEX idx_order_items_order_id ON order_items (order_id);

-- ==========================================
-- ORDER STATUS HISTORY (AUDIT TRAIL)
-- ==========================================
CREATE TABLE order_status_history
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT      NOT NULL,
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    changed_at TIMESTAMP   NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_order_history
        FOREIGN KEY (order_id)
            REFERENCES orders (id)
            ON DELETE CASCADE
);

CREATE INDEX idx_status_history_order_id ON order_status_history (order_id);
