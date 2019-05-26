PROBLEM DESCRIPTION:
Your goal here is to design an API that allows for hero tracking, much like the Vue problem
You are to implement an API (for which the skeleton already exists) that has the following capabilities
- Get      : return a JSON representation of the hero with the name supplied
- Make     : create a superhero according to the JSON body supplied
- Calamity : a calamity of the supplied level requires heroes with an equivalent combined powerlevel to address it.
             Takes a calamity with powerlevel and at least 1 hero. On success return a 200 with json response indicating the calamity has been resolved.
             Otherwise return a response indicating that the heroes were not up to the task. Addressing a calamity adds 1 point of exhaustion.
- Rest     : recover 1 point of exhaustion
- Retire   : retire a superhero, someone may take up his name for the future passing on the title
- Kill     : a superhero has passed away, his name may not be taken up again.

On success all endpoints should return a status code 200.

If a hero reaches an exhaustion level of maxExhaustion then they die.

You are free to decide what your API endpoints should be called and what shape they should take. You can modify any code in this file however you'd like.

NOTE: you may want to install postman or another request generating software of your choosing to make testing easier. (api is running on localhost port 8081)

NOTE the second: the API is receiving asynchronous requests to manage our super friends. As such, your hero access should be thread-safe for writes.
Even if the operations are extremely lightweight we want to make our application scalable.

NOTE the third: There are many ways to make whatever package-level tracking you implement thread-safe, feel free to change anything about this file (without changing the functionality of the program) to do so.
i.e. add package-level maps, add functions that take the hero struct as reference, modify the hero struct, creating access control paradigms
I highly recommend looking into channels, mutexes, and other golang memory management patterns and pick whatever you're most comfortable with.
For mad props: a timeout on the memory management process which returns a resource not available.

Bonus: If you're having fun (this is by no means necessary) you can make the calamity hold the heroes up for a time and delay their unlocking in a go-routine
example:
go func(h *hero) {
    time.Sleep(calamityTime)
    // release lock on hero
}(heroPtr)