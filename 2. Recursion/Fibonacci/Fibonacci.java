class Fibonacci{
    public static void main(String[] args) {
        System.out.println(fun(4));
    }
    public static int fun(int n){
        if(n==0 || n==1){
            return n;
        }
        return fun(n-1)+fun(n-2);
    }
}