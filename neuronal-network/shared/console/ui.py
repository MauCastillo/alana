from shared.models.models import COINS
from shared.utils.utils import Utils
from neuronal_network.train import Neuronal
import os

TABLE_DATA = os.environ.get("TABLE_DATA", "market-collector")


class UI:

    def main(self):
        print("Wellcome to train the neuronal network \n")
        print("What type train, do you want? \n")
        print("1) Single currency: \n")
        print("2) All Currency: \n")

        try:
            option = input("You can answer between [1 - 2]")
            optionInt = int(option)
        except:
             print("An exception occurred")
             print("Try again! \n")
             self.main()

        match optionInt:
            case 1:
                print("Option 1")
                self.load_information_single()
            case 2:
                self.load_information_all()
            case _:
                print("Sorry, but the answer is wrong! try again! \n")
                self.main()

    def load_information_single(self):
        
        print("Choose the currency: \n")
        print("1) eth_usdt: \n")
        print("2) matic_usdt: \n")
        print("3) sol_usdt: \n")
        print("4) ada_usdt: \n")
        print("5) bnb_usdt: \n")

        try:
            choose = input("You can answer between [1 - 5]: ")
            chooseInt = int(choose)

        except:
             print("An exception occurred")
             print("Try again! \n")
             self.load_information_single()



        name_coin = ""
        if chooseInt in COINS:
            name_coin = COINS[chooseInt]

        utils = Utils(TABLE_DATA)
        data_training = utils.get_training_data_unique(name_coin)

        if len(data_training) == 0:
            print("Datos de entrenamiento Vacios !!!")
            return

        inputRaw = data_training["inputs"]
        target = data_training["target"]
        input_size = len(data_training["inputs"][0])

        print("______ Records Numbers %s _____\n", len(inputRaw))

        ai = Neuronal(input_size, inputRaw, target, name_coin)
        ai.Training()

    def load_information_all(self):
        print("Loading all coins... : \n")

        utils = Utils(TABLE_DATA)
        data_training = utils.get_training_data_all()

        if len(data_training) == 0:
            print("Datos de entrenamiento Vacios !!!")
            return

        inputRaw = data_training["inputs"]
        target = data_training["target"]
        input_size = len(data_training["inputs"][0])

        print("______ Records Numbers %s _____\n", len(inputRaw))

        ai = Neuronal(input_size, inputRaw, target, "all market")
        ai.Training()