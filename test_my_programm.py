from selenium import webdriver
import unittest
import time


class TestReg(unittest.TestCase):
    def test_reg1(self):
        link = "http://suninjuly.github.io/registration1.html"
        browser = webdriver.Chrome()
        browser.get(link)

        # Ваш код, который заполняет обязательные поля
        input1 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your first name']")
        input1.send_keys('rqw')
        input2 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your last name']")
        input2.send_keys('QWE')
        input3 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your email']")
        input3.send_keys('qwe@asd.ru')
        input4 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your phone:']")
        input4.send_keys('ewq')
        input5 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your address:']")
        input5.send_keys('qwer')
        # Отправляем заполненную форму
        button = browser.find_element_by_css_selector("button.btn")
        button.click()

        # Проверяем, что смогли зарегистрироваться
        # ждем загрузки страницы
        time.sleep(1)

        # находим элемент, содержащий текст
        welcome_text_elt = browser.find_element_by_tag_name("h1")
        # записываем в переменную welcome_text текст из элемента welcome_text_elt
        welcome_text = welcome_text_elt.text

        # с помощью assert проверяем, что ожидаемый текст совпадает
        self.assertEqual("Congratulations! You have successfully registered!", welcome_text, "нет такого текста")
        browser.close()
        browser.quit()

    def test_reg2(self):
        link = "http://suninjuly.github.io/registration2.html"
        browser = webdriver.Chrome()
        browser.get(link)

        # Ваш код, который заполняет обязательные поля
        input1 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your first name']")
        input1.send_keys('rqw')
        input2 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your last name']")
        input2.send_keys('QWE')
        input3 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your email']")
        input3.send_keys('qwe@asd.ru')
        input4 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your phone:']")
        input4.send_keys('ewq')
        input5 = browser.find_element_by_xpath(
            "//input[@placeholder='Input your address:']")
        input5.send_keys('qwer')
        # Отправляем заполненную форму
        button = browser.find_element_by_css_selector("button.btn")
        button.click()

        # Проверяем, что смогли зарегистрироваться
        # ждем загрузки страницы
        time.sleep(1)

        # находим элемент, содержащий текст
        welcome_text_elt = browser.find_element_by_tag_name("h1")
        # записываем в переменную welcome_text текст из элемента welcome_text_elt
        welcome_text = welcome_text_elt.text

        # с помощью assert проверяем, что ожидаемый текст совпадает
        self.assertEqual("Congratulations! You have successfully registered!", welcome_text, "нет такого текста")
        # закрываем браузер после всех манипуляций
        browser.close()
        browser.quit()

if __name__ == "__main__":
    unittest.main()