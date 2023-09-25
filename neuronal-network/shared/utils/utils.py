import numpy as np
from shared.database.dynamodb import Dynamodb
from shared.models.models import COINS_INDEX

class Utils:
    FLOAT_PRECISION = 2
    input_data = []
    target_data = []

    def __init__(self, table):
        self.Table = table
    
    def parse_float(self, input_string):
        result_float = round(float(input_string), self.FLOAT_PRECISION)
        return result_float

    def market_array(self, record):
        outputmarket = []

        for kline in record['L']:
            outputmarket = np.append(
                outputmarket, self.parse_float(kline["N"]))

        return outputmarket

    def unique_array(self, record):
        output = np.array([], dtype=np.float)

        coin_name = record["name"]["S"]

        if coin_name in  COINS_INDEX:
            output = np.append(output, COINS_INDEX[coin_name])
        else:
            output = np.append(output, 0)

        value = self.parse_float(record["fear_greed_previous_1_month"]["N"])
        output = np.append(output, value)

        output = np.append(output, self.parse_float(
            record["fear_greed_previous_1_year"]["N"]))

        output = np.append(output, self.parse_float(
            record["fear_greed_previous_close"]["N"]))

        output = np.append(output, self.parse_float(
            record["fear_greed_score"]["N"]))

        output = np.append(output, self.parse_float(
            record["junk_bond_demand_score"]["N"]))

        output = np.append(output, self.parse_float(
            record["market_momentum_sp125_score"]["N"]))

        output = np.append(output, self.parse_float(
            record["market_momentum_sp500_score"]["N"]))

        output = np.append(output, self.parse_float(
            record["relative_strenght_index"]["N"]))

        output = np.append(output, self.parse_float(
            record["safe_haven_demand_score"]["N"]))

        output = np.append(output, self.parse_float(
            record["stochastic_oscillator_d"]["N"]))

        output = np.append(output, self.parse_float(
            record["stochastic_oscillator_k"]["N"]))

        output = np.append(output, self.parse_float(
            record["price_buy"]["N"]))

        output = np.append(output, self.market_array(record["market_info"]))

        output = np.append(output, self.market_array(
            record["market_info_btc"]))

        self.input_data.append(output)
        self.target_data = np.append(
            self.target_data, self.string_bool_to_int(record["status"]['BOOL']))

        return output
    
    def string_bool_to_int(self, value):
        if value:
            return 1
        
        return 0
    

    def get_training_data_unique(self, coin):
        database = Dynamodb(self.Table, coin)
        data_rows = database.read_table_raw()

        for record in data_rows:
            self.unique_array(record)

        print("input: ", len(self.input_data))
        print("input: ", self.input_data[5])
        print("target: ", self.target_data)

        return {"input": self.input_data, "target": self.target_data}
