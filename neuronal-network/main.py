import sys
import numpy as np
from database.dynamodb import Dynamodb
from neuronal_network.training import Neuronal


def load_information_dynamodb():

    database = Dynamodb("market-collector", "eth_usdt")
    data_training = database.get_data()

    if len(data_training) == 0:
        print("Datos de entrenamiento Vacios!!!!")
        return



    input = data_training["input"]
    target = data_training["target"]
    input_size = len(data_training["input"][0])
    


    ai = Neuronal(input_size, input, target)
    ai.Training()

def main():
   load_information_dynamodb()


if __name__ == "__main__":
    sys.exit(main())
