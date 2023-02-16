from telethon.sync import TelegramClient
import os, sys, json, time

class Eye():
    def __init__(self, api_hash, api_id, namefilelog, phone, delay_sec):
        self.ErrorFindUser = "UserNotFound"
        self.api_hash = api_hash 
        self.api_id = api_id
        self.namefile = namefilelog
        self.target_user = phone
        self.status_user = ""
        self.delay_sec = delay_sec
        self.symsplit = "&"
        self.client = TelegramClient("newSessionEye", api_id, api_hash)
        self.client.start()
        # warn init line
        for dialog in self.client.iter_dialogs():
            break
        self.check_file()


    def check_file(self):
        if os.path.isfile(self.namefile) == False:
            with open(self.namefile, "w+") as file:
                file.read() 

    def get_status(self, user_name):
        try:
            acc = str(self.client.get_entity(user_name).status).split("(")[0]
            return f"{acc}"
        except Exception as e:
            print(e)
            return f"{self.ErrorFindUser}"


    def write_file_line(self, status):
        with open(self.namefile, "a") as file:
            file.write(time.asctime() + self.symsplit + status + "\n")


    def old_status_from_file(self):
        with open(self.namefile, "r") as file:
            data = file.read().strip()
            if data == "":
                return self. ErrorFindUser
            data = data.split("\n")[-1]
            data = data.split(self.symsplit)[-1]
            return data
        
    def run(self):
        result = self.get_status(self.target_user)
        if self.old_status_from_file() != result:
            self.write_file_line(result)



if len(sys.argv) != 2:	
    print("Error name")
    sys.exit(0)

target_user = sys.argv[1]

with open('config/config_py.json', 'r') as f:
  config = json.load(f)
  

namefile = f"file_eye_{target_user}.log"
eye = Eye(config["api_hash"], config["api_id"], namefile, target_user, config["delay_sec"])
eye.run()
