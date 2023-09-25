import tensorflow as tf
import numpy as np
import matplotlib.pyplot as plt


class Neuronal:
    HIDDE_LAYERs_SIZE = 3
    EPOCHS = 20
    MODEL_NAME = "model_backup.h5"

    def __init__(self, input_size, input, output):
        self.input_data = input
        self.target_data = output
        self.input_size = input_size

    def Training(self):
        input_formatter = np.array(self.input_data, dtype=float)
        output_formatter = np.array(self.target_data, dtype=float)

        print(">>> input_formatter: ", input_formatter[0])

        layers = []

        intput_layer = tf.keras.layers.Dense(
            units=self.input_size, input_shape=[self.input_size])
        layers.append(intput_layer)

        for _ in range(self.HIDDE_LAYERs_SIZE):
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
        model.save(self.MODEL_NAME)

        print("Staring training")
        history = model.fit(input_formatter, output_formatter,
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
