class Power{
    public static void main(String[] args) {
        System.out.println(powFun(3, 2));
    }
    public static int powFun(int n, int k){
        if(k==0){
            return 1;
        }
        return n*powFun(n, k-1);
    }
}