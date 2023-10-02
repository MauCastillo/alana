import tensorflow as tf
import numpy as np
import matplotlib.pyplot as plt


class Neuronal:
    HIDDE_LAYERS_SIZE = 6
    EPOCHS = 20
    MODEL_NAME = "storage_models/model_backup_%s.h5"

    def __init__(self, input_size, input, target, coin_name):
        self.input_data = input
        self.target_data = target
        self.input_size = input_size
        self.coin_name = coin_name

    def Training(self):
        input_formatter = np.array(self.input_data, dtype=float)
        target_formatter = np.array(self.target_data, dtype=float)

        print(">>> target: ", target_formatter)

        layers = []

        intput_layer = tf.keras.layers.Dense(
            units=self.input_size, input_shape=[self.input_size])
        layers.append(intput_layer)

        for _ in range(self.HIDDE_LAYERS_SIZE):
            hidde = tf.keras.layers.Dense(units=self.input_size, activation='relu')
            layers.append(hidde)

       

        output_layer = tf.keras.layers.Dense(units=1, activation='linear')
        layers.append(output_layer)

        model = tf.keras.Sequential(layers)

        model.compile(
            optimizer=tf.keras.optimizers.Adam(0.1),
            loss='mean_squared_error'
        )

        # Guardar el Modelo
        file_name = self.MODEL_NAME % self.coin_name
        model.save(file_name)

        print("Staring training")
        history = model.fit(input_formatter, target_formatter,
                            epochs=self.EPOCHS, verbose=False)
        print("Finished the training!")

        plt.xlabel("# Epoca")
        plt.ylabel("Magnitud de perdida")

        plt.plot(history.history["loss"])
        plt.show()

    def prediction(self, test_value):
        model = tf.keras.models.load_model('farehert.h5')
        result = model.predict([test_value])
        print("Resultado", result)
