import math

def number_of_rectangle(m, n):
	res = 0
	num_mult = m * n
	max_num = max(m, n)


	for i in range(1, n + 1):
		for k in range(1, m + 1):
			print(i, k)
			print(math.ceil(num_mult / (i * k)))
			res += math.ceil(num_mult / (i * k))

	# for num in range(1, max_num + 1):
	# 	if num <= max_num:
	# 		print(math.floor(num_mult / num))
	# 		res += math.floor(num_mult / num)
		# if num <= n:
		# 	print(math.floor(num_mult / num))
		# 	res += math.floor(num_mult / num)


	# for i in range(1, max_num+1):
	# 	print(i)
	# 	res +=i
	return res


print(number_of_rectangle(2, 3)) # 18