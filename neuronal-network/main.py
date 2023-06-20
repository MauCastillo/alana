import tensorflow as tf
import numpy as np
import matplotlib.pyplot as plt



def converCelsiusToFahrenheit(celsius):
    return celsius * 1.8 + 32


newCelsius = []
newFahrenheit = []
for i in range(500):
    newCelsius.append(i)
    newFahrenheit.append(converCelsiusToFahrenheit(i))
    #output = "C = %s \n F = %s" % (i, converCelsiusToFahrenheit(i))
    #print(output)



num_samples = 1000
input_data = np.random.rand(num_samples, 1)
#input_data = np.random.randint(100,size=10)
for value in input_data:
    print("[%s]" % (value))
addVector = np.vectorize(converCelsiusToFahrenheit)
target_data = addVector(input_data)
for value in target_data:
    print(">>> %s => <<<" % (value))

#target_data = np.sum(input_data, axis=1)

celsius = np.array(input_data, dtype=float)
fahrenheit = np.array(target_data, dtype=float)

#layer = tf.keras.layers.Dense(units=1, input_shape=[1])
layer = tf.keras.layers.Dense(units=10, input_shape=[1])
hidde = tf.keras.layers.Dense(units=10)
hidde2 = tf.keras.layers.Dense(units=10)
output = tf.keras.layers.Dense(units=1)

model = tf.keras.Sequential([layer, hidde,hidde2, output])

model.compile(
    optimizer=tf.keras.optimizers.Adam(0.1),
    loss='mean_squared_error'
)

print("Staring training")
history = model.fit(celsius,fahrenheit,epochs=1000, verbose=False)
print("Finished the training!")

plt.xlabel("# Epoca")
plt.ylabel("Magnitud de perdida")

plt.plot(history.history["loss"])
resultado = model.predict([[100], [0.0]])
print("Resultado",resultado)
#print(layer.get_weights())
plt.show() 