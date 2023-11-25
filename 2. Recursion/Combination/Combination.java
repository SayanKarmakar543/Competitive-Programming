class Combination{
    public static void main(String[] args) {
        System.out.println(com(3,1));
    }
    static int com(int n, int r){
        if(n==r || r==0){
            return 1;
        }
        return com(n-1, r)+ com(n-1, r-1);
    }
}