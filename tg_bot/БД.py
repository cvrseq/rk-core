import sqlite3
import datetime

DB_PATH = "users.db"
def init_db():
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()
    cur.execute('''
        CREATE TABLE IF NOT EXISTS users (
            user_id INTEGER PRIMARY KEY,
            username TEXT,
            balance INTEGER DEFAULT 0,
            sub_end TEXT DEFAULT NULL
        )
    ''')
    conn.commit()
    conn.close()

def add_user(user_id, username):
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()
    cur.execute("INSERT OR IGNORE INTO users (user_id, username) VALUES (?, ?)", (user_id, username))
    conn.commit()
    conn.close()

def get_user(user_id):
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()
    cur.execute("SELECT balance, sub_end FROM users WHERE user_id = ?", (user_id,))
    user_data = cur.fetchone()
    conn.close()
    return user_data # Вернет баланс + срок подписки


def add_subscription(user_id, days):
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()
# Проверка, есть ли активная подписка
    cur.execute("SELECT sub_end FROM users WHERE user_id = ?", (user_id,))
    result = cur.fetchone()

    now = datetime.datetime.now()

    if result and result[0] and result[0] != "None":
        try:
            current_end = datetime.datetime.strptime(result[0], "%Y-%m-%d %H:%M")
            if current_end < now:
                current_end = now
        except ValueError:
            current_end = now
    else:
        current_end = now

    new_end = current_end + datetime.timedelta(days=days)
    new_end_str = new_end.strftime("%Y-%m-%d %H:%M")

# Обновляем запись в базе
    cur.execute("UPDATE users SET sub_end = ? WHERE user_id = ?", (new_end_str, user_id))
    conn.commit()
    conn.close()
# Возвращаем новую дату окончания
    return new_end_str
# Функция подсчета пользователей
def get_users_count():
    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()
    cur.execute("SELECT COUNT(user_id) FROM users")
    count = cur.fetchone()[0]
    conn.close()
    return count