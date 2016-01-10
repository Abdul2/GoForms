
### Introduction  http://abdul2.github.io/goforms-public/

This project will introduce an architecture to build a form driven GO application. The idea is that all form driven web apps consist of: 

1. Display 
2. Data
3. logic 
4. Database

The architecture is basically based on the fact the the logic part  (GO func)

1. recieves the data
2. changes it to inform the next step
3. does some house keeping 
3. writes to database

`http req --> server --> router --> func handler(logic....) --> db`


### Motivation 

I wrote a form driven web app in Go to help a Dubia based business  (a friend's business) to manage their  property 
portfolio and learned one or two things about the use of GO interfaces, structs and templates to be able to reduce the amount of code to minimum.I will start with the architecture, write pseudo code first before writing actual code (in the first version i worked in 
pilot mode and compromised on elegance and good design to get something out of the door). 

### Sharing 

I will put out most of my code in the public domain but some of it is considered private. i will run two repos and the bulk of the work will be in the public repo.

### How will it work

it is simple 

- Form is displayed/rendered 
- Form sends struct to logic 
    - struct types are
        - business specific (in my case Asset and Unit) 
        - form specific. Form specific structs drives the dynamic elements of the form and it is the mechanism by which the current logic tells the next piece of logic what to do. the logic manipulates the form structs fields. 
        - errors. to record error messages and display them back to users
   
- logic receives the structs and operate on their contents and depending on the stages it does one or all of the 
following

1. validates the data
2. changes the data
3. writes data to database
4. loads up template and marshal data to it
        
so all what we need is :

1. define data structure  - Go structs 
2. define html templates  - Go driven html templates
3. write logic to operate on data  - Go func
4. figure out how we persist data - Go sql.DB
5. write code to render forms - GO html/template
    
### Basic principles     
   
1. shall only write once 
2. shall use interfaces 
3. shall document

### Show me the money (code)

.... will do but but let's do some thinking first