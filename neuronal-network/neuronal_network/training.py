import tensorflow as tf
import numpy as np
import matplotlib.pyplot as plt


class Neuronal:
    LAYER_SIZE = 132
    EPOCHS = 50
    MODEL_NAME = "model_backup.h5"

    def __init__(self, input_size, input, output):
        self.input_data = input
        self.target_data = output
        self.input_size = input_size

    def Training(self):
        input_formatter = np.array(self.input_data, dtype=float)
        output_formatter = np.array(self.target_data, dtype=float)

        print(">>> input_formatter: ",input_formatter[0])

        layer = tf.keras.layers.Dense(
            units=self.LAYER_SIZE, input_shape=[self.input_size])
        hidde = tf.keras.layers.Dense(units=self.LAYER_SIZE, activation='relu')
        hidde2 = tf.keras.layers.Dense(units=self.LAYER_SIZE, activation='relu')
        hidde3 = tf.keras.layers.Dense(units=self.LAYER_SIZE, activation='relu')
        hidde4 = tf.keras.layers.Dense(units=self.LAYER_SIZE, activation='relu')
        hidde5 = tf.keras.layers.Dense(units=self.LAYER_SIZE, activation='relu')
        output = tf.keras.layers.Dense(units=1)

        model = tf.keras.Sequential([layer, hidde, hidde2, hidde3, hidde4, hidde5,  output])

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
