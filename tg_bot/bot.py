import asyncio
import os
import datetime
from aiogram import Bot, Dispatcher, types, F
from aiogram.filters import CommandStart, Command
from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton, CallbackQuery, BufferedInputFile
import logging
from dotenv import load_dotenv
import БД

load_dotenv()
TOKEN = os.getenv('BOT_TOKEN')
ADMIN_ID = int(os.getenv('ADMIN_ID'))

dp = Dispatcher()

# Функция меню
def get_main_keyboard():
    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="🛒 Купить VPN", callback_data="buy_vpn")],
        [
            InlineKeyboardButton(text="👤 Мой профиль", callback_data="profile"),
            InlineKeyboardButton(text="🆘 Помощь", callback_data="help")
        ]
    ])
    return kb

def get_tariffs_keyboard():
    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="1 месяц - 150₽", callback_data="buy_1m")],
        [InlineKeyboardButton(text="3 месяца - 400₽", callback_data="buy_3m")],
        [InlineKeyboardButton(text="⬅️ Назад", callback_data="back_to_main")]
    ])
    return kb

# Реализация кнопок через декораторы
@dp.message(CommandStart())
async def cmd_start(message: types.Message):
    # Записываем пользователя в БД
    БД.add_user(message.from_user.id, message.from_user.username)
    await message.answer(
        f"Привет, {message.from_user.first_name}! 🛡\n\nЯ запомнил тебя в системе💋! Выбери действие:",
        reply_markup=get_main_keyboard()
    )

@dp.callback_query(F.data == "buy_vpn")
async def process_buy_button(callback: CallbackQuery):
    await callback.message.edit_text(
        text="Выбери подходящий тарифный план:",
        reply_markup=get_tariffs_keyboard()
    )
    await callback.answer()

@dp.callback_query(F.data == "profile")
async def process_profile_button(callback: CallbackQuery):
    user_data = БД.get_user(callback.from_user.id)
    balance, sub_end = user_data if user_data else (0, None)
    # Проверяем, активна ли подписка прямо сейчас
    is_active = False
    if sub_end and sub_end != "None":
        end_date = datetime.datetime.strptime(sub_end, "%Y-%m-%d %H:%M")
        if end_date > datetime.datetime.now():
            is_active = True

    status_text = "✅ Активна" if is_active else "❌ Неактивна"
    display_date = sub_end if sub_end else "Нет данных"

    text = (
        f"👤 <b>Мой профиль</b>\n\n"
        f"🆔 Ваш ID: <code>{callback.from_user.id}</code>\n"
        f"📊 Статус: <b>{status_text}</b>\n"
        f"⏳ Срок до: <b>{display_date}</b>\n"
        f"💰 Баланс: {balance}₽"
    )
    buttons = []
    # Если подписка активна - даем кнопку на скачивание ключа
    if is_active:
        buttons.append([InlineKeyboardButton(text="🔑 Получить ключ (AWG)", callback_data="get_my_key")])
    buttons.append([InlineKeyboardButton(text="⬅️ Назад", callback_data="back_to_main")])
    kb = InlineKeyboardMarkup(inline_keyboard=buttons)

    # Проверяем, есть ли у сообщения обычный текст
    if callback.message.text:
        # Если есть текст - просто редактируем его
        await callback.message.edit_text(text, parse_mode="HTML", reply_markup=kb)
    else:
        # Если текста нет - удаляем его и шлем новое сообщение
        await callback.message.delete()
        await callback.message.answer(text, parse_mode="HTML", reply_markup=kb)
    await callback.answer()

@dp.callback_query(F.data == "help")
async def process_help_button(callback: CallbackQuery):
    text = (
        "❓ <b>Центр поддержки</b>\n\n"
        "1️⃣ <b>Как подключиться?</b>\n"
        "Скачайте приложение AmneziaWG, нажмите 'Добавить туннель' и выберите файл, который прислал бот.\n\n"
        "2️⃣ <b>Не работает интернет?</b>\n"
        "Попробуйте переключить авиарежим или сменить сеть (Wi-Fi/4G).\n\n"
        "3️⃣ <b>Связь с оператором:</b>\n"
        "Если у вас возникла проблема с оплатой или доступом, напишите нашему администратору."
    )

    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="🆘 Написать админу", url="https://t.me/alexandertolcachev")],
        [InlineKeyboardButton(text="⬅️ Назад", callback_data="back_to_main")]
    ])

    await callback.message.edit_text(text, parse_mode="HTML", reply_markup=kb)
    await callback.answer()

@dp.callback_query(F.data.startswith("buy_") & (F.data != "buy_vpn"))
async def process_payment_selection(callback: CallbackQuery):
    action = callback.data
    if action == "buy_1m":
        price = 150
        duration = "1 месяц"
    elif action == "buy_3m":
        price = 400
        duration = "3 месяца"
    else:
        return

    text = (
        f"🧾 <b>Оформление заказа</b>\n\n"
        f"📦 Тариф: <b>VPN на {duration}</b>\n"
        f"💵 К оплате: <b>{price}₽</b>\n\n"
        f"<i>Здесь будет настоящая ссылка на оплату (например, YooMoney или CryptoBot).</i>"
    )

    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="💳 Оплатить", url="https://yoomoney.ru")],
        [InlineKeyboardButton(text="🔄 Я оплатил (Проверить)", callback_data=f"check_pay_{action}")],
        [InlineKeyboardButton(text="⬅️ Отмена", callback_data="buy_vpn")]  # Возвращает к выбору тарифов
    ])

    await callback.message.edit_text(text, parse_mode="HTML", reply_markup=kb)
    await callback.answer()

# ВАЖНО!
# Заглушка проверки оплаты. Будет реализовано через Микро-API, но можно через ssh, если уверен, что сделаешь из него крепость
@dp.callback_query(F.data.startswith("check_pay_"))
async def check_payment(callback: CallbackQuery):
    action = callback.data.replace("check_pay_", "")
    if action == "buy_1m":
        days = 30
    elif action == "buy_3m":
        days = 90
    else:
        days = 0
    # Имитация загрузки
    await callback.message.edit_text("⏳ <i>Генерируем конфигурацию AmnesiaWG... Это займет пару секунд.</i>",
                                     parse_mode="HTML")
    # Имитация ожидание ответа от сервера (API)
    await asyncio.sleep(2)
    # Запись в БД
    new_date = БД.add_subscription(callback.from_user.id, days)
    # Имитация ответа от Микро-API (Якобы конфиг AmnesiaWG)
    fake_awg_config = f"""[Interface]
PrivateKey = aBcDeFgHiJkLmNoPqRsTuVwXyZ123456789=
Address = 10.8.0.2/24
DNS = 1.1.1.1
Jc = 120
Jmin = 23
Jmax = 911
S1 = 57
S2 = 33
H1 = 1
H2 = 2
H3 = 3
H4 = 4

[Peer]
PublicKey = ZYXwVuTsRqPoNmLkJiHgFeDcBa987654321=
Endpoint = 192.168.1.1:51820
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25"""
    # Преобразование текста в настоящий файл .conf
    config_file = BufferedInputFile(
        fake_awg_config.encode('utf-8'),
        filename=f"awg_vpn_{callback.from_user.id}.conf"
    )
    # Кнопка возврата
    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="👤 В профиль", callback_data="profile")]
    ])
    # Отправка файла пользователю
    await callback.message.answer_document(
        document=config_file,
        caption=f"✅ <b>Оплата успешно найдена!</b>\n\n🎉 Подписка продлена до: <b>{new_date}</b>\n\nСкачай прикрепленный файл конфигурации и импортируй его в приложение <b>AmneziaWG</b> на своем устройстве.",
        parse_mode="HTML",
        reply_markup=kb
    )
    await callback.message.delete()
    await callback.answer()

@dp.callback_query(F.data == "get_my_key")
async def send_existing_key(callback: CallbackQuery):
    # Уведомление пользователя, что процесс пошел
    await callback.message.edit_text("⏳ <i>Подготавливаем ваш персональный ключ...</i>", parse_mode="HTML")
    # Имитация задержки "обращения к серверу"
    await asyncio.sleep(1)
    # Якобы конфиг AmnesiaWG
    fake_awg_config = f"""[Interface]
PrivateKey = aBcDeFgHiJkLmNoPqRsTuVwXyZ123456789=
Address = 10.8.0.2/24
DNS = 1.1.1.1
Jc = 120
Jmin = 23
Jmax = 911
S1 = 57
S2 = 33
H1 = 1
H2 = 2
H3 = 3
H4 = 4

[Peer]
PublicKey = ZYXwVuTsRqPoNmLkJiHgFeDcBa987654321=
Endpoint = 192.168.1.1:51820
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25"""
    # Превращаем текст конфигурации в полноценный файл .conf
    config_file = BufferedInputFile(
        fake_awg_config.encode('utf-8'),
        filename=f"awg_vpn_{callback.from_user.id}.conf"
    )
    # Делаем кнопку для возврата обратно в профиль
    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="⬅️ Назад в профиль", callback_data="profile")]
    ])

    # Отправляем сам файл с подписью и кнопкой возврата
    await callback.message.answer_document(
        document=config_file,
        caption="🔑 <b>Ваш ключ доступа</b>\n\nИмпортируйте этот файл в приложение <b>AmneziaWG</b>.",
        parse_mode="HTML",
        reply_markup=kb
    )
    await callback.message.delete()
    await callback.answer()

@dp.callback_query(F.data == "back_to_main")
async def process_back_button(callback: CallbackQuery):
    await callback.message.edit_text(
        text=f"Привет, {callback.from_user.first_name}! 🛡\n\nВыбери нужное действие в меню ниже:",
        reply_markup=get_main_keyboard()
    )
    await callback.answer()

# АДМИН ПАНЕЛЬ, НАЧАЛО
def get_admin_keyboard():
    kb = InlineKeyboardMarkup(inline_keyboard=[
        [InlineKeyboardButton(text="📊 Статистика", callback_data="admin_stats")],
        [InlineKeyboardButton(text="📢 Рассылка (в разработке)", callback_data="admin_broadcast")]
    ])
    return kb

# Секретная команда /admin
@dp.message(Command("admin"))
async def cmd_admin(message: types.Message):
    # Проверка является ли пользователь админом
    if message.from_user.id == ADMIN_ID:
        await message.answer("🛠 <b>Панель администратора</b>", parse_mode="HTML", reply_markup=get_admin_keyboard())
    else:
        # Если пишет обычный юзер, бот делает вид, что такой команды нет
        pass

# Стата для админа
@dp.callback_query(F.data == "admin_stats")
async def process_admin_stats(callback: CallbackQuery):
    if callback.from_user.id != ADMIN_ID:
        return
    users_count = БД.get_users_count()
    text = (
        f"📊 <b>Статистика проекта:</b>\n\n"
        f"👥 Всего пользователей: <b>{users_count}</b>\n"
        f"<i>(Здесь мы потом добавим подсчет заработанных денег)</i>"
    )
    await callback.message.edit_text(text, parse_mode="HTML", reply_markup=get_admin_keyboard())
    await callback.answer()
# АДМИН ПАНЕЛЬ КОНЕЦ
async def main():
    logging.basicConfig(level=logging.INFO)
    БД.init_db()
    bot = Bot(token=TOKEN)
    print("Бот успешно запущен!")
    await dp.start_polling(bot)

if __name__ == '__main__':
    asyncio.run(main())