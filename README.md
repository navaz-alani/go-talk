# Go-Talk

Go-Talk is an ephemeral chat web-app. It is intended for untraceable
chat between two parties. Once a message has been delivered, it only
exists on the sender's and recipient's devices for the time that it is
visible. 

In order to ensure transport security, I plan to implement end-to-end
encryption for payloads being sent through the core service. With the
current hosting, security is as good as SSL over a TCP connection. 
While this can prevent packet sniffing by an intruder, packets are 
still readable as they transit through the backend's core service. 
Therefore, the backend could be considered the weakest link (security
-wise) in the whole chain.
